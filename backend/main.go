package main

import (
	"log"

	"app.com/backend/server"
	"github.com/subosito/gotenv"
)

func init() {
	// loads values from .env into the system
	if err := gotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	server := server.NewServer()

	server.Run()
}
