package requesthandlers

import (
	"net/http"
	"strings"

	"github.com/cfindlayisme/wmb/env"
	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/cfindlayisme/wmb/model"
	"github.com/gin-gonic/gin"
)

func validateMessage(msg model.IncomingMessage, c *gin.Context) bool {
	if strings.Contains(msg.Message, "\n") || strings.Contains(msg.Message, "\r") {
		c.String(http.StatusBadRequest, "Message cannot contain newline characters")
		return false
	}
	if env.GetWebhookPassword() != msg.Password {
		c.String(http.StatusUnauthorized, "Invalid password")
		return false
	}
	return true
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
