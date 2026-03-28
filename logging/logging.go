package logging

import (
	"log"

	"github.com/cfindlayisme/wmb/env"
)

func DebugLog(v ...interface{}) {
	if env.GetDebug() {
		log.Println(v...)
	}
}

func DebugLogf(format string, v ...interface{}) {
	if env.GetDebug() {
		log.Printf(format, v...)
	}
}
