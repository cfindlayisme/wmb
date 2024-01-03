package ircclient

import "github.com/cfindlayisme/wmb/model"

func FormatMessage(msg model.IncomingMessage) string {
	ircMessage := msg.Message

	colourPrefix := ""
	colourSufffix := "\x03"
	// https://modern.ircdocs.horse/formatting.html
	switch msg.ColourCode {
	case 0:
		colourPrefix = "\x0300"
	case 1:
		colourPrefix = "\x0301"
	case 2:
		colourPrefix = "\x0302"
	case 3:
		colourPrefix = "\x0303"
	case 4:
		colourPrefix = "\x0304"
	case 5:
		colourPrefix = "\x0305"
	case 6:
		colourPrefix = "\x0306"
	case 7:
		colourPrefix = "\x0307"
	case 8:
		colourPrefix = "\x0308"
	case 9:
		colourPrefix = "\x0309"
	case 10:
		colourPrefix = "\x0310"
	case 11:
		colourPrefix = "\x0311"
	case 12:
		colourPrefix = "\x0312"
	case 13:
		colourPrefix = "\x0313"
	case 14:
		colourPrefix = "\x0314"
	case 15:
		colourPrefix = "\x0315"
	}

	if colourPrefix != "" {
		return colourPrefix + ircMessage + colourSufffix
	}

	return ircMessage
}
