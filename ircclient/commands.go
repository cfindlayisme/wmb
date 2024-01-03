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

func SendMessage(conn net.Conn, target string, message string) error {
	_, err := fmt.Fprintf(conn, "PRIVMSG "+target+" :"+message+"\r\n")
	return err
}

func SetUser(conn net.Conn) error {
	_, err := fmt.Fprintf(conn, "USER wmb 0 * :Webhook message bot\r\n")
	return err
}
