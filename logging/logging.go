package logging

import (
	"log"

	"github.com/cfindlayisme/wmb/env"
)

const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
	colorBoldRed = "\033[1;31m"
)

func Debug(v ...interface{}) {
	if env.GetDebug() {
		log.Println(append([]interface{}{colorCyan + "[DEBUG]" + colorReset}, v...)...)
	}
}

func Debugf(format string, v ...interface{}) {
	if env.GetDebug() {
		log.Printf(colorCyan+"[DEBUG]"+colorReset+" "+format, v...)
	}
}

func Info(v ...interface{}) {
	log.Println(append([]interface{}{colorGreen + "[INFO]" + colorReset}, v...)...)
}

func Infof(format string, v ...interface{}) {
	log.Printf(colorGreen+"[INFO]"+colorReset+" "+format, v...)
}

func Warn(v ...interface{}) {
	log.Println(append([]interface{}{colorYellow + "[WARN]" + colorReset}, v...)...)
}

func Warnf(format string, v ...interface{}) {
	log.Printf(colorYellow+"[WARN]"+colorReset+" "+format, v...)
}

func Error(v ...interface{}) {
	log.Println(append([]interface{}{colorRed + "[ERROR]" + colorReset}, v...)...)
}

func Errorf(format string, v ...interface{}) {
	log.Printf(colorRed+"[ERROR]"+colorReset+" "+format, v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(append([]interface{}{colorBoldRed + "[FATAL]" + colorReset}, v...)...)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(colorBoldRed+"[FATAL]"+colorReset+" "+format, v...)
}
