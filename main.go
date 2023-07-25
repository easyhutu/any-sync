package main

import (
	"github.com/easyhutu/any-sync/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := app.NewAnySyncServer()
	server.Run()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	server.Stop()
}
