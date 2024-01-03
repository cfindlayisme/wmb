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
	ColourCode int8
	// 0 default colour
	// 1 green
	// 2 red
	// 3 blue
	// 4 yellow
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
