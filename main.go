package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/cfindlayisme/wmb/env"
	"github.com/gin-gonic/gin"
)

type IncomingMessage struct {
	Message    string
	Password   string
	ColourCode int8 // https://modern.ircdocs.horse/formatting.html
}

func main() {
	conn, err := net.Dial("tcp", env.GetServer())
	if err != nil {
		fmt.Println("Failed to connect to IRC server:", err)
		os.Exit(1)
	}

	fmt.Fprintf(conn, "NICK "+env.GetNick()+"\r\n")
	fmt.Fprintf(conn, "USER wmb 0 * :Webhook message bot\r\n")
	fmt.Fprintf(conn, "JOIN "+env.GetChannel()+"\r\n")

	router := gin.Default()
	listenAddress := "0.0.0.0:8080"

	go func() {
		router.POST("/message", func(c *gin.Context) {
			var msg IncomingMessage

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

				_, err := fmt.Fprintf(conn, "PRIVMSG "+env.GetChannel()+" :"+ircMessage+"\r\n")

				if err != nil {
					c.String(http.StatusInternalServerError, "Failed to send message to IRC server")
				}
				c.String(http.StatusOK, "Message sent")
			}

		})
		router.Run(listenAddress)
	}()

	for {
		message, err := readMessage(conn)
		if err != nil {
			fmt.Println("Failed to read message:", err)
			break
		}

		fmt.Println("Received message:", message)

		if strings.HasPrefix(message, "PING") {
			pongMessage := strings.Replace(message, "PING", "PONG", 1)
			fmt.Fprintf(conn, pongMessage+"\r\n")
			fmt.Println("Sent message:", pongMessage)
		}
	}

	conn.Close()
}

func readMessage(conn net.Conn) (string, error) {
	buffer := make([]byte, 512)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}
