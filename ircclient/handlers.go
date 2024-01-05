package ircclient

import (
	"fmt"
	"strings"
)

func returnPong(message string) {
	pongMessage := strings.Replace(message, "PING", "PONG", 1)
	fmt.Fprintf(ircConnection, pongMessage+"\r\n")
	fmt.Println("PONG returned to server PING")
}

func processPrivmsg(words []string) {
	// Extract the channel and the message from the PRIVMSG command
	// The format of a PRIVMSG command is: :nick!user@host PRIVMSG #channel :message
	if len(words) < 4 {
		fmt.Println("Invalid PRIVMSG command:", strings.Join(words, " "))
		return
	}

	// Extract the nick, user, and host from the first word
	// The format of the first word is :nick!user@host
	prefix := strings.TrimPrefix(words[0], ":")
	prefixParts := strings.SplitN(prefix, "!", 2)
	if len(prefixParts) < 2 {
		fmt.Println("Invalid prefix in PRIVMSG command:", prefix)
		return
	}

	nick := prefixParts[0]

	userHostParts := strings.SplitN(prefixParts[1], "@", 2)
	if len(userHostParts) < 2 {
		fmt.Println("Invalid user@host in PRIVMSG command:", prefixParts[1])
		return
	}

	user := userHostParts[0]
	host := userHostParts[1]

	channel := words[2]
	msg := strings.Join(words[3:], " ") // The message can contain spaces, so join all remaining words

	fmt.Printf("Received PRIVMSG from %s!%s@%s to %s: %s\n", nick, user, host, channel, msg)
}
