package model

import "time"

type IncomingMessage struct {
	Message    string
	Password   string
	ColourCode *int8 // https://modern.ircdocs.horse/formatting.html
	Broadcast  *bool
}

type DirectedIncomingMessage struct {
	Target          string
	IncomingMessage IncomingMessage
}

type DirectedOutgoingMessage struct {
	Target    string
	Message   string
	IRCUser   IrcUser
	Timestamp time.Time
}

type PrivmsgSubscription struct {
	Target   string
	URL      string
	Password string
}

type IrcUser struct {
	Nick string
	User string
	Host string
}
