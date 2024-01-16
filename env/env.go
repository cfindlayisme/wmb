package env

import (
	"os"
	"strings"
)

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

func GetOtherChannels() []string {
	channels := os.Getenv("OTHER_IRC_CHANNELS")
	if channels == "" {
		return nil
	}
	return strings.Split(channels, ",")
}

func GetNick() string {
	nick := os.Getenv("IRC_NICK")
	if nick == "" {
		nick = "wmb"
	}
	return nick
}

func GetWebhookPassword() string {
	return os.Getenv("PASSWORD")
}

func GetNickservPassword() string {
	return os.Getenv("NICKSERV_PASSWORD")
}

func GetDatabaseFile() string {
	dbfile := os.Getenv("DBFILE")
	if dbfile == "" {
		dbfile = "wmb.db"
	}
	return os.Getenv("DBFILE")
}

func GetListenPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
