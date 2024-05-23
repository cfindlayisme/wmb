package requesthandlers

import (
	"net/http"
	"strings"

	"github.com/cfindlayisme/wmb/env"
	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/cfindlayisme/wmb/model"
	"github.com/cfindlayisme/wmb/webhook"
	"github.com/gin-gonic/gin"
)

func validateMessage(msg model.IncomingMessage, c *gin.Context) bool {
	if strings.Contains(msg.Message, "\n") || strings.Contains(msg.Message, "\r") {
		c.String(http.StatusBadRequest, "Message cannot contain newline characters")
		return false
	}
	validatePassword(msg.Password, c)

	return true
}

func validatePassword(password string, c *gin.Context) bool {
	if env.GetWebhookPassword() == password {
		return true
	}
	c.String(http.StatusUnauthorized, "Invalid password")

	return false
}

func PostMessage(c *gin.Context) {
	var msg model.IncomingMessage

	if err := c.BindJSON(&msg); err != nil {
		c.String(http.StatusBadRequest, "Invalid JSON in request body")
		return
	}

	if !validateMessage(msg, c) {
		return
	}

	err := ircclient.SendMessage(env.GetChannel(), ircclient.FormatMessage(msg))

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to send message to IRC server")
	}
	c.String(http.StatusOK, "Message sent")

}

func PostDirectedMessage(c *gin.Context) {
	var dmsg model.DirectedIncomingMessage
	var msg model.IncomingMessage

	if err := c.BindJSON(&dmsg); err != nil {
		c.String(http.StatusBadRequest, "Invalid JSON in request body")
		return
	}

	msg = dmsg.IncomingMessage

	if !validateMessage(msg, c) {
		return
	}

	err := ircclient.SendMessage(dmsg.Target, ircclient.FormatMessage(msg))

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to send message to IRC server")
	}
	c.String(http.StatusOK, "Message sent")

}

func QueryMessage(c *gin.Context) {
	var msg model.IncomingMessage

	if err := c.ShouldBindQuery(&msg); err != nil {
		c.String(http.StatusBadRequest, "Invalid query parameters")
		return
	}

	if !validateMessage(msg, c) {
		return
	}

	err := ircclient.SendMessage(env.GetChannel(), ircclient.FormatMessage(msg))

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to send message to IRC server")
	}
	c.String(http.StatusOK, "Message sent")

}

func PostSubscribePrivmsg(c *gin.Context) {
	var subscription model.PrivmsgSubscription

	if err := c.BindJSON(&subscription); err != nil {
		c.String(http.StatusBadRequest, "Invalid query parameters")
		return
	}

	if !validatePassword(subscription.Password, c) {
		return
	}

	success := webhook.SubscribePrivmsg(subscription.Target, subscription.URL)

	if success {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Subscription successful",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "Subscription failed",
		})
	}
}
