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
	Target  string `json:"target"`
	Message string `json:"message"`
}
