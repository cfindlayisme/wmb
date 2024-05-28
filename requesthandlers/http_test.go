package requesthandlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/cfindlayisme/wmb/model"
	"github.com/cfindlayisme/wmb/pointers"
	"github.com/cfindlayisme/wmb/requesthandlers"
	"github.com/stretchr/testify/assert"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
)

func TestEndpointsPasswordProtection(t *testing.T) {
	os.Setenv("PASSWORD", "correct_password")
	gin.SetMode(gin.TestMode)

	// Mock the SendMessage function
	monkey.Patch(ircclient.SendMessage, func(conn net.Conn, target string, message string) error {
		return nil
	})

	defer monkey.UnpatchAll()

	tests := []struct {
		name     string
		method   string
		endpoint string
		body     interface{}
	}{
		{"PostMessage", "POST", "/message", &model.IncomingMessage{Password: "wrong_password"}},
		{"QueryMessage", "GET", "/message", &model.IncomingMessage{Password: "wrong_password"}},
		{"PostDirectedMessage", "POST", "/directedMessage", &model.DirectedIncomingMessage{IncomingMessage: model.IncomingMessage{Password: "wrong_password"}}},
		{"PostSubscribePrivmsg", "POST", "/subscribe/message", &model.PrivmsgSubscription{Password: "wrong_password"}},
		{"PostUnsubscribePrivmsg", "POST", "/unsubscribe/message", &model.PrivmsgSubscription{Password: "wrong_password"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()

			switch tt.endpoint {
			case "/message":
				if tt.method == "POST" {
					router.POST(tt.endpoint, requesthandlers.PostMessage)
				} else if tt.method == "GET" {
					router.GET(tt.endpoint, requesthandlers.QueryMessage)
				}
			case "/directedMessage":
				router.POST(tt.endpoint, requesthandlers.PostDirectedMessage)
			case "/subscribe/message":
				router.POST(tt.endpoint, requesthandlers.PostSubscribePrivmsg)
			case "/unsubscribe/message":
				router.POST(tt.endpoint, requesthandlers.PostUnsubscribePrivmsg)
			}

			var req *http.Request
			if tt.method == "GET" {
				req, _ = http.NewRequest(tt.method, tt.endpoint+"?Password=wrong_password", nil)
			} else {
				body, _ := json.Marshal(tt.body)
				req, _ = http.NewRequest(tt.method, tt.endpoint, bytes.NewBuffer(body))
			}

			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			if status := resp.Code; status != http.StatusUnauthorized {
				t.Errorf("Handler returned wrong status code: got %v want %v",
					status, http.StatusUnauthorized)
			}
		})
	}
}

func TestSendBroadcastMessage(t *testing.T) {
	// Set up environment variables
	os.Setenv("PASSWORD", "correct_password")
	os.Setenv("IRC_CHANNEL", "channel1")
	os.Setenv("OTHER_IRC_CHANNELS", "channel2,channel3")

	// Create a Gin router
	router := gin.Default()
	router.POST("/message", requesthandlers.PostMessage)
	router.GET("/message", requesthandlers.QueryMessage)

	// Test cases where Broadcast is true, false or not set
	testCases := []struct {
		name      string
		broadcast *bool
		channel1  bool
		channel2  bool
		channel3  bool
	}{
		{"BroadcastTrue", pointers.BoolPtr(true), true, true, true},
		{"BroadcastFalse", pointers.BoolPtr(false), true, false, false},
		{"BroadcastNotSet", nil, true, false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a map to track which channels messages are sent to
			channelMap := make(map[string]bool)

			// Patch the SendMessage function
			monkey.Patch(ircclient.SendMessage, func(conn net.Conn, channel string, message string) error {
				channelMap[channel] = true
				return nil
			})

			// Create a test request for PostMessage
			postBody, _ := json.Marshal(&model.IncomingMessage{Password: "correct_password", Broadcast: tc.broadcast})
			postReq, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(postBody))
			postResp := httptest.NewRecorder()
			router.ServeHTTP(postResp, postReq)

			// Create a test request for QueryMessage
			queryReq, _ := http.NewRequest("GET", fmt.Sprintf("/message?Password=correct_password&Broadcast=%v", tc.broadcast), nil)
			queryResp := httptest.NewRecorder()
			router.ServeHTTP(queryResp, queryReq)

			// Check that messages were sent to the correct channels
			assert.Equal(t, tc.channel1, channelMap["channel1"])
			assert.Equal(t, tc.channel2, channelMap["channel2"])
			assert.Equal(t, tc.channel3, channelMap["channel3"])

			// Unpatch the SendMessage function after each test case
			monkey.Unpatch(ircclient.SendMessage)
		})
	}
}

func TestPostMessageSuccess(t *testing.T) {
	// Set up environment variables
	os.Setenv("PASSWORD", "correct_password")
	os.Setenv("IRC_CHANNEL", "channel1")

	// Create a Gin router
	router := gin.Default()
	router.POST("/message", requesthandlers.PostMessage)

	// Variables to track if the message was sent to the correct channel and its contents
	var messageSent bool
	var sentMessage string

	// Patch the SendMessage function
	monkey.Patch(ircclient.SendMessage, func(conn net.Conn, channel string, message string) error {
		if channel == "channel1" {
			messageSent = true
			sentMessage = message
		}
		return nil
	})

	// Create a test request for PostMessage
	testMessage := "Test message"
	incomingMessage := model.IncomingMessage{Message: testMessage, Password: "correct_password"}
	postBody, _ := json.Marshal(&incomingMessage)
	postReq, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(postBody))
	postResp := httptest.NewRecorder()
	router.ServeHTTP(postResp, postReq)

	// Check that the message was sent to the correct channel
	assert.True(t, messageSent)

	// Check that the message contents match
	assert.Equal(t, ircclient.FormatMessage(incomingMessage), sentMessage)

	// Check that the response status is OK
	assert.Equal(t, http.StatusOK, postResp.Code)

	// Unpatch the SendMessage function after the test
	monkey.Unpatch(ircclient.SendMessage)
}
