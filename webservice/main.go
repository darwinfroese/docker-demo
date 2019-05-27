package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/darwinfroese/docker-demo/webservice/server"
)

const shutdownWaitDuration time.Duration = time.Second * 15

func main() {
	srv := server.InitializeRouter()

	log.Println("Launching server...")

	go func() {
		log.Println("Server running...")
		if err := srv.ListenAndServe(); err != nil {
			log.Print("Server crashed! ")
			log.Println(err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), shutdownWaitDuration)
	defer cancel()

	log.Println("Terminating web server...")
	srv.Shutdown(ctx)
	os.Exit(0)
}
