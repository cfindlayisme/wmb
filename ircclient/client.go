package ircclient

import (
	"log"
	"net"
	"strings"

	"github.com/cfindlayisme/wmb/env"
)

var ircConnection net.Conn

func Connect(server string) error {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		return err
	}

	setConnection(conn)

	initialize()

	return nil
}

func initialize() {
	SetNick(env.GetNick())
	SetUser()
	if env.GetNickservPassword() != "" {
		SendMessage("NickServ", "IDENTIFY "+env.GetNickservPassword())
	}

	if env.GetSelfMode() != "" {
		SetMode(env.GetNick(), env.GetSelfMode())
	}

	// Join our primary channel
	JoinChannel(env.GetChannel())

	// Also join our non-primary channels
	for _, channel := range env.GetOtherChannels() {
		JoinChannel(channel)
	}
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
			log.Println("Failed to read message on TCP buffer:", err)
			break
		}

		message = strings.TrimSuffix(message, "\n")
		words := strings.Split(message, " ")

		if strings.HasPrefix(message, "PING") {
			ReturnPong(ircConnection, message)
		} else if len(words) >= 2 && words[1] == "PRIVMSG" {
			processPrivmsg(words)
		} else {
			log.Println("Raw unprocessed message:", message)
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
