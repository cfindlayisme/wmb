package ircclient

import (
	"fmt"
	"net"
	"strings"

	"github.com/cfindlayisme/wmb/logging"
	"github.com/cfindlayisme/wmb/model"
	"github.com/cfindlayisme/wmb/webhook"
)

func ReturnPong(connection net.Conn, message string) {
	pongMessage := strings.Replace(message, "PING", "PONG", 1)
	fmt.Fprintf(connection, pongMessage+"\r\n")
	logging.Debug("PONG returned to server PING")
}

func processPrivmsg(words []string) {
	// Extract the channel and the message from the PRIVMSG command
	// The format of a PRIVMSG command is: :nick!user@host PRIVMSG #channel :message
	if len(words) < 4 {
		logging.Warn("Invalid PRIVMSG command:", strings.Join(words, " "))
		return
	}

	// Extract the nick, user, and host from the first word
	// The format of the first word is :nick!user@host
	prefix := strings.TrimPrefix(words[0], ":")
	prefixParts := strings.SplitN(prefix, "!", 2)
	if len(prefixParts) < 2 {
		logging.Warn("Invalid prefix in PRIVMSG command:", prefix)
		return
	}

	nick := prefixParts[0]

	userHostParts := strings.SplitN(prefixParts[1], "@", 2)
	if len(userHostParts) < 2 {
		logging.Warn("Invalid user@host in PRIVMSG command:", prefixParts[1])
		return
	}

	user := userHostParts[0]
	host := userHostParts[1]

	channel := words[2]
	msg := strings.Join(words[3:], " ") // The message can contain spaces, so join all remaining words
	msg = strings.TrimPrefix(msg, ":")  // Remove leading :
	msg = strings.TrimSuffix(msg, "\r") // Remove trailing \r

	// Check if the message is a CTCP request
	if strings.HasPrefix(msg, "\x01") && strings.HasSuffix(msg, "\x01") {
		processCTCP(nick, user, host, msg[1:len(msg)-1]) // Remove wrapping \x01
		return
	}

	logging.Infof("Received PRIVMSG from %s!%s@%s to %s: %s", nick, user, host, channel, msg)

	ircuser := model.IrcUser{
		Nick: nick,
		User: user,
		Host: host,
	}
	webhook.SendPrivmsgWebhook(channel, msg, ircuser)
}

func processCTCP(nick, user, host, msg string) {
	// CTCP requests are formatted as \x01COMMAND data\x01
	msg = strings.TrimSuffix(msg, "\x01") // Remove trailing \x01
	parts := strings.SplitN(msg, " ", 2)  // Split command and optional data
	command := parts[0]

	switch command {

	case "VERSION":
		logging.Infof("Received CTCP VERSION request from %s!%s@%s", nick, user, host)
		SendCTCPReply(IrcConnection, nick, "VERSION", "wmb - github.com/cfindlayisme/wmb")
	default:
		logging.Warnf("Unknown CTCP command '%s' from %s!%s@%s", command, nick, user, host)

	}
}
