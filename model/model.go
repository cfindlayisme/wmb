package model

type IncomingMessage struct {
	Message    string
	Password   string
	ColourCode int8 // https://modern.ircdocs.horse/formatting.html
}

type DirectedIncomingMessage struct {
	Target          string
	IncomingMessage IncomingMessage
}

type DirectedOutgoingMessage struct {
	Target  string
	Message string
	IRCUser IrcUser
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
