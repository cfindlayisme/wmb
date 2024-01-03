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
	ircclient.JoinChannel(env.GetChannel())

	router := gin.Default()
	listenAddress := "0.0.0.0:8080"

	// Web server needs to be launched as a goroutine so that it doesn't block
	go func() {
		router.POST("/message", requesthandlers.PostMessage)

		err := router.Run(listenAddress)

		if err != nil {
			fmt.Println("Failed to start webserver:", err)
			os.Exit(1)
		}
	}()

	ircclient.Loop()

	ircclient.Disconnect()
}
