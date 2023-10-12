package main

import (
	. "elephant_carpaccio/domain"
	http_server "elephant_carpaccio/http-server"
	"log"
	"net/http"
)

func main() {
	renderer, _ := http_server.NewRenderer()
	game := NewGame()
	server := http_server.NewBoardServer(renderer, game)
	log.Fatal(http.ListenAndServe("localhost:3000", server))
}
