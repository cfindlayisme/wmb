package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/cfindlayisme/wmb/env"
	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/cfindlayisme/wmb/model"
	"github.com/gin-gonic/gin"
)

func main() {
	err := ircclient.Connect(env.GetServer())
	if err != nil {
		fmt.Println("Failed to connect to IRC server:", err)
		os.Exit(1)
	}

	ircclient.SetNick(env.GetNick())
	ircclient.SetUser()
	ircclient.JoinChannel(env.GetChannel())

	router := gin.Default()
	listenAddress := "0.0.0.0:8080"

	go func() {
		router.POST("/message", func(c *gin.Context) {
			var msg model.IncomingMessage

			if err := c.BindJSON(&msg); err != nil {
				c.String(http.StatusBadRequest, "Invalid JSON in request body")
				return
			}

			if env.GetWebhookPassword() != msg.Password {
				c.String(http.StatusUnauthorized, "Invalid password")
			} else {
				ircMessage := msg.Message
				// Strip newlines to prevent chaining of commands, ie, QUIT to the end
				ircMessage = strings.ReplaceAll(ircMessage, "\n", "")
				ircMessage = strings.ReplaceAll(ircMessage, "\r", "")

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

		})

		err := router.Run(listenAddress)

		if err != nil {
			fmt.Println("Failed to start webserver:", err)
			os.Exit(1)
		}
	}()

	for {
		message, err := readMessage(ircclient.GetConnection())
		if err != nil {
			fmt.Println("Failed to read message:", err)
			break
		}

		fmt.Println("Received message:", message)

		if strings.HasPrefix(message, "PING") {
			ircclient.ReturnPong(message)
		}
	}

	ircclient.Disconnect()
}

func readMessage(conn net.Conn) (string, error) {
	buffer := make([]byte, 512)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}
