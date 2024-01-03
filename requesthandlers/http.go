package requesthandlers

import (
	"net/http"
	"strings"

	"github.com/cfindlayisme/wmb/env"
	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/cfindlayisme/wmb/model"
	"github.com/gin-gonic/gin"
)

func validateMessage(msg model.IncomingMessage, c *gin.Context) {
	if strings.Contains(msg.Message, "\n") || strings.Contains(msg.Message, "\r") {
		c.String(http.StatusBadRequest, "Message cannot contain newline characters")
	}
}

func PostMessage(c *gin.Context) {
	var msg model.IncomingMessage

	if err := c.BindJSON(&msg); err != nil {
		c.String(http.StatusBadRequest, "Invalid JSON in request body")
		return
	}

	validateMessage(msg, c)

	if env.GetWebhookPassword() != msg.Password {
		c.String(http.StatusUnauthorized, "Invalid password")
	} else {
		ircMessage := msg.Message

		colourPrefix := ""
		colourSufffix := "\x03"
		// https://modern.ircdocs.horse/formatting.html
		switch msg.ColourCode {
		case 0:
			colourPrefix = "\x0300"
		case 1:
			colourPrefix = "\x0301"
		case 2:
			colourPrefix = "\x0302"
		case 3:
			colourPrefix = "\x0303"
		case 4:
			colourPrefix = "\x0304"
		case 5:
			colourPrefix = "\x0305"
		case 6:
			colourPrefix = "\x0306"
		case 7:
			colourPrefix = "\x0307"
		case 8:
			colourPrefix = "\x0308"
		case 9:
			colourPrefix = "\x0309"
		case 10:
			colourPrefix = "\x0310"
		case 11:
			colourPrefix = "\x0311"
		case 12:
			colourPrefix = "\x0312"
		case 13:
			colourPrefix = "\x0313"
		case 14:
			colourPrefix = "\x0314"
		case 15:
			colourPrefix = "\x0315"
		}

		if colourPrefix != "" {
			ircMessage = colourPrefix + ircMessage + colourSufffix
		}

		err := ircclient.SendMessage(env.GetChannel(), ircMessage)

		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to send message to IRC server")
		}
		c.String(http.StatusOK, "Message sent")
	}

}
