package env

import "os"

func GetServer() string {
	return os.Getenv("IRC_SERVER")
}
