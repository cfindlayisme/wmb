package ircclient

import (
	"fmt"
	"net"
)

func SetNick(conn net.Conn, nick string) error {
	_, err := fmt.Fprintf(conn, "NICK "+nick+"\r\n")
	return err
}

func JoinChannel(conn net.Conn, channel string) error {
	_, err := fmt.Fprintf(conn, "JOIN "+channel+"\r\n")
	return err
}

func PartChannel(conn net.Conn, channel string) error {
	_, err := fmt.Fprintf(conn, "PART "+channel+"\r\n")
	return err
}

func SetMode(conn net.Conn, channel string, mode string) error {
	_, err := fmt.Fprintf(conn, "MODE "+channel+" "+mode+"\r\n")
	return err
}

func SetTopic(conn net.Conn, channel string, topic string) error {
	cleanTopic := CleanMessage(topic)
	_, err := fmt.Fprintf(conn, "TOPIC "+channel+" "+cleanTopic+"\r\n")
	return err
}

func InviteUser(conn net.Conn, nick string, channel string) error {
	_, err := fmt.Fprintf(conn, "INVITE "+nick+" "+channel+"\r\n")
	return err
}

func KickUser(conn net.Conn, nick string, channel string, message string) error {
	if message == "" {
		message = "Kicked"
	}
	_, err := fmt.Fprintf(conn, "KICK "+channel+" "+nick+" :"+message+"\r\n")
	return err
}

func Quote(conn net.Conn, command string) error {
	_, err := fmt.Fprintf(conn, command+"\r\n")
	return err
}

func SendMessage(conn net.Conn, target string, message string) error {
	ircMessage := CleanMessage(message)

	_, err := fmt.Fprintf(conn, "PRIVMSG "+target+" :"+ircMessage+"\r\n")
	return err
}

func SendNotice(conn net.Conn, target string, message string) error {
	ircMessage := CleanMessage(message)

	_, err := fmt.Fprintf(conn, "NOTICE "+target+" :"+ircMessage+"\r\n")
	return err
}

func SetUser(conn net.Conn) error {
	_, err := fmt.Fprintf(conn, "USER wmb 0 * :Webhook message bot\r\r\n")
	return err
}

func SendQuit(conn net.Conn, quitMessage string) error {
	_, err := fmt.Fprintf(conn, "QUIT :%s\r\n", CleanMessage(quitMessage))
	return err
}

func SendCTCPReply(conn net.Conn, target, command, response string) error {
	ctcpMessage := fmt.Sprintf("\x01%s %s\x01", command, CleanMessage(response))
	_, err := fmt.Fprintf(conn, "NOTICE %s :%s\r\n", target, ctcpMessage)
	return err
}
