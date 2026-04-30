package ircclient

import (
	"fmt"
	"log"
	"net"

	goutilsstrings "github.com/cfindlayisme/go-utils/strings"
)

func SetNick(conn net.Conn, nick string) error {
	_, err := fmt.Fprintf(conn, "NICK %s\r\n", nick)
	log.Println("Set nick to ", nick)
	return err
}

func JoinChannel(conn net.Conn, channel string) error {
	_, err := fmt.Fprintf(conn, "JOIN %s\r\n", channel)
	log.Println("JOIN command sent for channel: ", channel)
	return err
}

func PartChannel(conn net.Conn, channel string) error {
	_, err := fmt.Fprintf(conn, "PART %s\r\n", channel)
	log.Println("PART command sent for channel: ", channel)
	return err
}

func SetMode(conn net.Conn, channel string, mode string) error {
	_, err := fmt.Fprintf(conn, "MODE %s %s\r\n", channel, mode)
	log.Println("MODE command sent for target: ", channel, " with mode: ", mode)
	return err
}

func SetTopic(conn net.Conn, channel string, topic string) error {
	cleanTopic := goutilsstrings.StripNewlines(topic)
	_, err := fmt.Fprintf(conn, "TOPIC %s %s\r\n", channel, cleanTopic)
	log.Println("TOPIC command sent for channel: ", channel, " with topic: ", cleanTopic)
	return err
}

func InviteUser(conn net.Conn, nick string, channel string) error {
	_, err := fmt.Fprintf(conn, "INVITE %s %s\r\n", nick, channel)
	log.Println("INVITE command sent for nick: ", nick, " to channel: ", channel)
	return err
}

func KickUser(conn net.Conn, nick string, channel string, message string) error {
	if message == "" {
		message = "Kicked"
	}
	_, err := fmt.Fprintf(conn, "KICK %s %s :%s\r\n", channel, nick, message)
	log.Println("KICK command sent for nick: ", nick, " from channel: ", channel, " with message: ", message)
	return err
}

func Quote(conn net.Conn, command string) error {
	_, err := fmt.Fprintf(conn, "%s\r\n", command)
	log.Println("Sent raw command: ", command)
	return err
}

func SendMessage(conn net.Conn, target string, message string) error {
	ircMessage := goutilsstrings.StripNewlines(message)

	_, err := fmt.Fprintf(conn, "PRIVMSG %s :%s\r\n", target, ircMessage)
	log.Println("Sent message to ", target, ": ", ircMessage)
	return err
}

func SendNotice(conn net.Conn, target string, message string) error {
	ircMessage := goutilsstrings.StripNewlines(message)

	_, err := fmt.Fprintf(conn, "NOTICE %s :%s\r\n", target, ircMessage)
	log.Println("Sent notice to ", target, ": ", ircMessage)
	return err
}

func SetUser(conn net.Conn) error {
	_, err := fmt.Fprintf(conn, "USER wmb 0 * :Webhook message bot\r\n")
	log.Println("Sent USER command")
	return err
}

func SendQuit(conn net.Conn, quitMessage string) error {
	_, err := fmt.Fprintf(conn, "QUIT :%s\r\n", goutilsstrings.StripNewlines(quitMessage))
	log.Println("Sent QUIT command")
	return err
}

func SendCTCPReply(conn net.Conn, target, command, response string) error {
	ctcpMessage := fmt.Sprintf("\x01%s %s\x01", command, goutilsstrings.StripNewlines(response))
	_, err := fmt.Fprintf(conn, "NOTICE %s :%s\r\n", target, ctcpMessage)
	return err
}
