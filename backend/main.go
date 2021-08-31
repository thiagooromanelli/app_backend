package main

import (
	"app.com/backend/server"
)

func main() {
	server := server.NewServer()

	server.Run()
}
