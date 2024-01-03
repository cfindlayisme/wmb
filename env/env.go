package env

import "os"

func GetServer() string {
	return os.Getenv("IRC_SERVER")
}

func GetChannel() string {
	channel := os.Getenv("IRC_CHANNEL")
	if channel == "" {
		channel = "#wmb"
	}
	return channel
}

func GetNick() string {
	nick := os.Getenv("IRC_NICK")
	if nick == "" {
		nick = "wmb"
	}
	return nick
}
