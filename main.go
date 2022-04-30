package main

import (
	"log"

	"github.com/SHA65536/arcade/server"
)

func main() {
	server, err := server.MakeServer("127.0.0.1", 9992)
	if err != nil {
		log.Fatal(err)
	}
	server.Run()
}
