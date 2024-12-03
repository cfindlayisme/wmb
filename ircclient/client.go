package ircclient

import (
	"bufio"
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
}

func initializePostConnect() {
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
	SendQuit(IrcConnection, "Disconnecting!")

	return IrcConnection.Close()
}

func setConnection(conn net.Conn) {
	IrcConnection = conn
}

func Loop() {
	isPostConnect := false // Tracks if the connection is ready for commands

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
		} else if len(words) >= 2 && words[1] == "001" {
			// 001 Welcome message
			isPostConnect = true
		} else if len(words) >= 2 && (words[1] == "376" || words[1] == "422") {
			// 376: End of MOTD, 422: No MOTD
			isPostConnect = true
		} else {
			log.Println("Raw unprocessed message:", message)
		}

		if isPostConnect {
			initializePostConnect()
			log.Println("Connected to IRC server - doing post-connect routine")
			isPostConnect = false
		}
	}
}

func readMessage(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)

	message, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	message = strings.TrimRight(message, "\r\n")
	return message, nil
}
