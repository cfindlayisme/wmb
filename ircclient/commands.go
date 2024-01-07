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

func PartChannel(channel string) error {
	_, err := fmt.Fprintf(ircConnection, "PART "+channel+"\r\n")
	return err
}

func SetMode(channel string, mode string) error {
	_, err := fmt.Fprintf(ircConnection, "MODE "+channel+" "+mode+"\r\n")
	return err
}

func SetTopic(channel string, topic string) error {
	cleanTopic := cleanMessage(topic)
	_, err := fmt.Fprintf(ircConnection, "TOPIC "+channel+" "+cleanTopic+"\r\n")
	return err
}

func InviteUser(nick string, channel string) error {
	_, err := fmt.Fprintf(ircConnection, "INVITE "+nick+" "+channel+"\r\n")
	return err
}

func Quote(command string) error {
	_, err := fmt.Fprintf(ircConnection, command+"\r\n")
	return err
}

func SendMessage(target string, message string) error {
	ircMessage := cleanMessage(message)

	_, err := fmt.Fprintf(ircConnection, "PRIVMSG "+target+" :"+ircMessage+"\r\n")
	return err
}

func SendNotice(target string, message string) error {
	ircMessage := cleanMessage(message)

	_, err := fmt.Fprintf(ircConnection, "NOTICE "+target+" :"+ircMessage+"\r\n")
	return err
}

func SetUser() error {
	_, err := fmt.Fprintf(ircConnection, "USER wmb 0 * :Webhook message bot\r\n")
	return err
}
