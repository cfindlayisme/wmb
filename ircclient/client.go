package ircclient

import (
	"fmt"
	"net"
	"strings"
)

var ircConnection net.Conn

func Connect(server string) error {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		return err
	}

	setConnection(conn)

	return nil
}

func Disconnect() error {
	return ircConnection.Close()
}

func setConnection(conn net.Conn) {
	ircConnection = conn
}

func Loop() {
	for {
		message, err := readMessage(ircConnection)
		if err != nil {
			fmt.Println("Failed to read message:", err)
			break
		}

		if strings.HasPrefix(message, "PING") {
			returnPong(message)
		}

		message = strings.TrimSuffix(message, "\n")
		words := strings.Split(message, " ")

		if len(words) >= 2 && words[1] == "PRIVMSG" {
			processPrivmsg(words)
		} else {
			fmt.Println("Raw unprocessed message:", message)
		}
	}
}

func readMessage(conn net.Conn) (string, error) {
	buffer := make([]byte, 512)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}
