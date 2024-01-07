package main

import (
	"fmt"
	"os"

	"github.com/cfindlayisme/wmb/env"
	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/cfindlayisme/wmb/requesthandlers"
	"github.com/gin-gonic/gin"
)

func main() {
	err := ircclient.Connect(env.GetServer())
	if err != nil {
		fmt.Println("Failed to connect to IRC server:", err)
		os.Exit(1)
	}

	ircclient.SetNick(env.GetNick())
	ircclient.SetUser()
	if env.GetNickservPassword() != "" {
		ircclient.SendMessage("NickServ", "IDENTIFY "+env.GetNickservPassword())
	}
	// Join our primary channel
	ircclient.JoinChannel(env.GetChannel())

	// Also join our non-primary channels
	for _, channel := range env.GetOtherChannels() {
		ircclient.JoinChannel(channel)
	}

	router := gin.Default()
	listenAddress := "0.0.0.0:8080"

	// Web server needs to be launched as a goroutine so that it doesn't block
	go func() {
		router.POST("/message", requesthandlers.PostMessage)
		router.GET("/message", requesthandlers.QueryMessage)
		router.POST("/directedMessage", requesthandlers.PostDirectedMessage)

		err := router.Run(listenAddress)

		if err != nil {
			fmt.Println("Failed to start webserver:", err)
			os.Exit(1)
		}
	}()

	ircclient.Loop()

	ircclient.Disconnect()
}
