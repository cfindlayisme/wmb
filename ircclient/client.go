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
	messageChannel := make(chan string)
	errorChannel := make(chan error)

	// Start a goroutine for reading messages
	go func() {
		readMessages(IrcConnection, messageChannel, errorChannel)
	}()

	isPostConnect := false

	for {
		select {
		case message, ok := <-messageChannel:
			if !ok {
				return // Exit the loop when the channel is closed
			}

			words := strings.Split(message, " ")

			if strings.HasPrefix(message, "PING") {
				ReturnPong(IrcConnection, message)
			} else if len(words) >= 2 && words[1] == "PRIVMSG" {
				processPrivmsg(words)
			} else if len(words) >= 2 && words[1] == "001" {
				isPostConnect = true
			} else if len(words) >= 2 && (words[1] == "376" || words[1] == "422") {
				isPostConnect = true
			} else {
				log.Println("Raw unprocessed message:", message)
			}

			if isPostConnect {
				initializePostConnect()
				log.Println("Connected to IRC server - doing post-connect routine")
				isPostConnect = false
			}

		case err := <-errorChannel:
			log.Println("Error reading from connection:", err)
			return
		}
	}
}

func readMessages(conn net.Conn, messageChannel chan<- string, errorChannel chan<- error) {
	reader := bufio.NewReaderSize(conn, 65536) // Increase buffer size for high-traffic servers

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			errorChannel <- err
			close(messageChannel) // Signal the loop to stop
			return
		}

		// Handle multiple lines in a single read
		messages := strings.Split(strings.TrimRight(message, "\r\n"), "\r\n")
		for _, msg := range messages {
			messageChannel <- msg
		}
	}
}
