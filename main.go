package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cfindlayisme/wmb/database"
	"github.com/cfindlayisme/wmb/env"
	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/cfindlayisme/wmb/router"
)

func main() {
	database.DB.Open(env.GetDatabaseFile())
	defer database.DB.Close()

	err := ircclient.Connect(env.GetServer())
	if err != nil {
		log.Fatalf("Failed to connect to IRC server: %s", err)
	}

	router := router.SetupRouter()
	listenAddress := "0.0.0.0:" + env.GetListenPort()

	// Web server needs to be launched as a goroutine so that it doesn't block
	go func() {
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
		database.DB.Close()
	}()

	ircclient.Loop()
}
