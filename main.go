package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cfindlayisme/wmb/database"
	"github.com/cfindlayisme/wmb/env"
	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/cfindlayisme/wmb/requesthandlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.DB.Open(env.GetDatabaseFile())
	defer database.DB.Close()

	err := ircclient.Connect(env.GetServer())
	if err != nil {
		log.Fatalf("Failed to connect to IRC server: %s", err)
	}

	router := gin.Default()
	listenAddress := "0.0.0.0:" + env.GetListenPort()

	// Web server needs to be launched as a goroutine so that it doesn't block
	go func() {
		router.POST("/message", requesthandlers.PostMessage)
		router.GET("/message", requesthandlers.QueryMessage)
		router.POST("/directedMessage", requesthandlers.PostDirectedMessage)
		router.POST("/subscribe/message", requesthandlers.PostSubscribePrivmsg)
		router.POST("/unsubscribe/message", requesthandlers.PostUnsubscribePrivmsg)

		err := router.Run(listenAddress)

		if err != nil {
			log.Fatalf("Failed to start webserver: %s", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Launch a goroutine to handle termination signals
	go func() {
		<-sigs

		ircclient.Disconnect()
	}()

	ircclient.Loop()
}
