package env

import "os"

func GetServer() string {
	return os.Getenv("IRC_SERVER")
}

func GetChannel() string {
	return os.Getenv("IRC_CHANNEL")
}

func GetNick() string {
	return os.Getenv("IRC_NICK")
}
