package main

import (
	. "elephant_carpaccio/domain"
	httpserver "elephant_carpaccio/http-server"
	"log"
	"net"
	"net/http"
)

func main() {
	game := NewGame()
	server := httpserver.NewBoardServer(httpserver.NewRenderer(net.InterfaceAddrs), game)
	log.Fatal(http.ListenAndServe(":3000", server))
}
