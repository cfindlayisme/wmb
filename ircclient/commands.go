package ircclient

import (
	"fmt"
)

func SetNick(nick string) error {
	_, err := fmt.Fprintf(ircConnection, "NICK "+nick+"\r\n")
	return err
}

func JoinChannel(channel string) error {
	_, err := fmt.Fprintf(ircConnection, "JOIN "+channel+"\r\n")
	return err
}

func SendMessage(target string, message string) error {
	_, err := fmt.Fprintf(ircConnection, "PRIVMSG "+target+" :"+message+"\r\n")
	return err
}

func SetUser() error {
	_, err := fmt.Fprintf(ircConnection, "USER wmb 0 * :Webhook message bot\r\n")
	return err
}
