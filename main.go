package main

import (
	. "elephant_carpaccio/domain"
	httpserver "elephant_carpaccio/http-server"
	"flag"
	"log"
	"net"
	"net/http"
)

func main() {
	var environment string
	flag.StringVar(&environment, "env", "PRD", "environment: DEV (development), PRD (production)")
	flag.Parse()

	println("options: ")
	flag.PrintDefaults()
	println("env: " + environment)

	game := NewGame()
	server := httpserver.NewBoardServer(game, net.InterfaceAddrs)

	if environment == "DEV" {
		log.Fatal(http.ListenAndServe("localhost:3000", server))
	} else {
		log.Fatal(http.ListenAndServe(":3000", server))
	}
}
