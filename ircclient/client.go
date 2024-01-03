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
func GetConnection() net.Conn {
	return ircConnection
}

func ReturnPong(message string) {
	pongMessage := strings.Replace(message, "PING", "PONG", 1)
	fmt.Fprintf(ircConnection, pongMessage+"\r\n")
	fmt.Println("PONG returned to server PING")
}
