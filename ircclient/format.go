package ircclient

import (
	"github.com/cfindlayisme/wmb/model"
)

func FormatMessage(msg model.IncomingMessage) string {
	ircMessage := msg.Message

	colourPrefix := ""
	colourSuffix := ""

	if msg.ColourCode == nil {
		return ircMessage
	}

	// https://modern.ircdocs.horse/formatting.html
	switch *msg.ColourCode {
	case 0:
		colourPrefix = "\x0300"
		colourSuffix = "\x03"
	case 1:
		colourPrefix = "\x0301"
		colourSuffix = "\x03"
	case 2:
		colourPrefix = "\x0302"
		colourSuffix = "\x03"
	case 3:
		colourPrefix = "\x0303"
		colourSuffix = "\x03"
	case 4:
		colourPrefix = "\x0304"
		colourSuffix = "\x03"
	case 5:
		colourPrefix = "\x0305"
		colourSuffix = "\x03"
	case 6:
		colourPrefix = "\x0306"
		colourSuffix = "\x03"
	case 7:
		colourPrefix = "\x0307"
		colourSuffix = "\x03"
	case 8:
		colourPrefix = "\x0308"
		colourSuffix = "\x03"
	case 9:
		colourPrefix = "\x0309"
		colourSuffix = "\x03"
	case 10:
		colourPrefix = "\x0310"
		colourSuffix = "\x03"
	case 11:
		colourPrefix = "\x0311"
		colourSuffix = "\x03"
	case 12:
		colourPrefix = "\x0312"
		colourSuffix = "\x03"
	case 13:
		colourPrefix = "\x0313"
		colourSuffix = "\x03"
	case 14:
		colourPrefix = "\x0314"
		colourSuffix = "\x03"
	case 15:
		colourPrefix = "\x0315"
		colourSuffix = "\x03"
	}

	if colourPrefix != "" {
		return colourPrefix + ircMessage + colourSuffix
	}

	return ircMessage
}
