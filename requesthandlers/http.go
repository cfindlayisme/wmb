package requesthandlers

import (
	"errors"
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

	return true
}

func validatePassword(password string) error {
	if env.GetWebhookPassword() == password {
		return nil
	}

	return errors.New("invalid password")
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

	if err := validatePassword(msg.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	var err error

	if msg.Broadcast != nil && *msg.Broadcast {
		channels := env.GetOtherChannels()
		channels = append(channels, env.GetChannel())
		for _, channel := range channels {
			err = ircclient.SendMessage(ircclient.IrcConnection, channel, ircclient.FormatMessage(msg))
			if err != nil {
				break
			}
		}
	} else {
		err = ircclient.SendMessage(ircclient.IrcConnection, env.GetChannel(), ircclient.FormatMessage(msg))
	}

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

	if err := validatePassword(msg.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	err := ircclient.SendMessage(ircclient.IrcConnection, dmsg.Target, ircclient.FormatMessage(msg))

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to send message to IRC server")
		return
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

	password := c.Query("password") // Extract password from query parameters
	if err := validatePassword(password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	err := ircclient.SendMessage(ircclient.IrcConnection, env.GetChannel(), ircclient.FormatMessage(msg))

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to send message to IRC server")
		return
	}
	c.String(http.StatusOK, "Message sent")
}

func PostSubscribePrivmsg(c *gin.Context) {
	var subscription model.PrivmsgSubscription

	if err := c.BindJSON(&subscription); err != nil {
		c.String(http.StatusBadRequest, "Invalid query parameters")
		return
	}

	if err := validatePassword(subscription.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
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

func PostUnsubscribePrivmsg(c *gin.Context) {
	var subscription model.PrivmsgSubscription

	if err := c.BindJSON(&subscription); err != nil {
		c.String(http.StatusBadRequest, "Invalid query parameters")
		return
	}

	if err := validatePassword(subscription.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	success := webhook.UnsubscribePrivmsg(subscription.Target, subscription.URL)

	if success {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Subscription removal successful",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "Subscription removal failed",
		})
	}
}
