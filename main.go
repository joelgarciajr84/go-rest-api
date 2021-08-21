package main

import (
	"log"

	"github.com/joelgarciajr84/go-rest-api/server"
)

func main() {
	log.Print("UP AND RUNNING")
	server := server.NewServer()

	server.Run()
}
