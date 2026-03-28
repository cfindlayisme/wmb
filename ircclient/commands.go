package ircclient

import (
	"fmt"
	"net"

	goutilsstrings "github.com/cfindlayisme/go-utils/strings"
	"github.com/cfindlayisme/wmb/logging"
)

func SetNick(conn net.Conn, nick string) error {
	_, err := fmt.Fprintf(conn, "NICK "+nick+"\r\n")
	logging.DebugLog("Set nick to ", nick)
	return err
}

func JoinChannel(conn net.Conn, channel string) error {
	_, err := fmt.Fprintf(conn, "JOIN "+channel+"\r\n")
	logging.DebugLog("JOIN command sent for channel: ", channel)
	return err
}

func PartChannel(conn net.Conn, channel string) error {
	_, err := fmt.Fprintf(conn, "PART "+channel+"\r\n")
	logging.DebugLog("PART command sent for channel: ", channel)
	return err
}

func SetMode(conn net.Conn, channel string, mode string) error {
	_, err := fmt.Fprintf(conn, "MODE "+channel+" "+mode+"\r\n")
	logging.DebugLog("MODE command sent for target: ", channel, " with mode: ", mode)
	return err
}

func SetTopic(conn net.Conn, channel string, topic string) error {
	cleanTopic := goutilsstrings.StripNewlines(topic)
	_, err := fmt.Fprintf(conn, "TOPIC "+channel+" "+cleanTopic+"\r\n")
	logging.DebugLog("TOPIC command sent for channel: ", channel, " with topic: ", cleanTopic)
	return err
}

func InviteUser(conn net.Conn, nick string, channel string) error {
	_, err := fmt.Fprintf(conn, "INVITE "+nick+" "+channel+"\r\n")
	logging.DebugLog("INVITE command sent for nick: ", nick, " to channel: ", channel)
	return err
}

func KickUser(conn net.Conn, nick string, channel string, message string) error {
	if message == "" {
		message = "Kicked"
	}
	_, err := fmt.Fprintf(conn, "KICK "+channel+" "+nick+" :"+message+"\r\n")
	logging.DebugLog("KICK command sent for nick: ", nick, " from channel: ", channel, " with message: ", message)
	return err
}

func Quote(conn net.Conn, command string) error {
	_, err := fmt.Fprintf(conn, command+"\r\n")
	logging.DebugLog("Sent raw command: ", command)
	return err
}

func SendMessage(conn net.Conn, target string, message string) error {
	ircMessage := goutilsstrings.StripNewlines(message)

	_, err := fmt.Fprintf(conn, "PRIVMSG "+target+" :"+ircMessage+"\r\n")
	logging.DebugLog("Sent message to ", target, ": ", ircMessage)
	return err
}

func SendNotice(conn net.Conn, target string, message string) error {
	ircMessage := goutilsstrings.StripNewlines(message)

	_, err := fmt.Fprintf(conn, "NOTICE "+target+" :"+ircMessage+"\r\n")
	logging.DebugLog("Sent notice to ", target, ": ", ircMessage)
	return err
}

func SetUser(conn net.Conn) error {
	_, err := fmt.Fprintf(conn, "USER wmb 0 * :Webhook message bot\r\n")
	logging.DebugLog("Sent USER command")
	return err
}

func SendQuit(conn net.Conn, quitMessage string) error {
	_, err := fmt.Fprintf(conn, "QUIT :%s\r\n", goutilsstrings.StripNewlines(quitMessage))
	logging.DebugLog("Sent QUIT command")
	return err
}

func SendCTCPReply(conn net.Conn, target, command, response string) error {
	ctcpMessage := fmt.Sprintf("\x01%s %s\x01", command, goutilsstrings.StripNewlines(response))
	_, err := fmt.Fprintf(conn, "NOTICE %s :%s\r\n", target, ctcpMessage)
	return err
}
