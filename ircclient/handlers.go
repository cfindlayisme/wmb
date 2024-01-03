package ircclient

import (
	"fmt"
	"strings"
)

func returnPong(message string) {
	pongMessage := strings.Replace(message, "PING", "PONG", 1)
	fmt.Fprintf(ircConnection, pongMessage+"\r\n")
	fmt.Println("PONG returned to server PING")
}
