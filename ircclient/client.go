package ircclient

import (
	"log"
	"net"
	"strings"

	"github.com/cfindlayisme/wmb/env"
)

var IrcConnection net.Conn

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
	SetNick(IrcConnection, env.GetNick())
	SetUser(IrcConnection)
	if env.GetNickservPassword() != "" {
		SendMessage(IrcConnection, "NickServ", "IDENTIFY "+env.GetNickservPassword())
	}

	if env.GetSelfMode() != "" {
		SetMode(IrcConnection, env.GetNick(), env.GetSelfMode())
	}

	// Join our primary channel
	JoinChannel(IrcConnection, env.GetChannel())

	// Also join our non-primary channels
	for _, channel := range env.GetOtherChannels() {
		JoinChannel(IrcConnection, channel)
	}
}

func Disconnect() error {
	return IrcConnection.Close()
}

func setConnection(conn net.Conn) {
	IrcConnection = conn
}

func Loop() {
	for {
		message, err := readMessage(IrcConnection)
		if err != nil {
			log.Println("Failed to read message on TCP buffer:", err)
			break
		}

		message = strings.TrimSuffix(message, "\n")
		words := strings.Split(message, " ")

		if strings.HasPrefix(message, "PING") {
			ReturnPong(IrcConnection, message)
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
