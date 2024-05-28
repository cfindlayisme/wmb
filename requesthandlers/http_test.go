package requesthandlers_test

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/cfindlayisme/wmb/model"
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

	// Create a map to track which channels messages are sent to
	channelMap := make(map[string]bool)

	// Patch the SendMessage function
	monkey.Patch(ircclient.SendMessage, func(conn net.Conn, channel string, message string) error {
		channelMap[channel] = true
		return nil
	})

	// Create a Gin router
	router := gin.Default()
	router.POST("/message", requesthandlers.PostMessage)
	router.GET("/message", requesthandlers.QueryMessage)

	// Create a test request for PostMessage
	broadcast := true
	postBody, _ := json.Marshal(&model.IncomingMessage{Password: "correct_password", Broadcast: &broadcast})
	postReq, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(postBody))
	postResp := httptest.NewRecorder()
	router.ServeHTTP(postResp, postReq)

	// Create a test request for QueryMessage
	queryReq, _ := http.NewRequest("GET", "/message?Password=correct_password&Broadcast=true", nil)
	queryResp := httptest.NewRecorder()
	router.ServeHTTP(queryResp, queryReq)

	// Check that messages were sent to all channels
	assert.True(t, channelMap["channel1"])
	assert.True(t, channelMap["channel2"])
	assert.True(t, channelMap["channel3"])
}
