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
		simulateGameForDev(game)
		log.Fatal(http.ListenAndServe("localhost:3000", server))
	} else {
		log.Fatal(http.ListenAndServe(":3000", server))
	}
}

func simulateGameForDev(game *Game) {
	game.Register("A Team")
	game.Register("The fantastic four")

	team1 := game.Teams()[0]
	team1.Done("EC-001", "EC-002", "EC-003")
	team1.CompleteIteration()
	team1.Done("EC-004", "EC-005")
	team1.CompleteIteration()
	team1.Done("EC-006", "EC-007", "EC-008")
	team1.CompleteIteration()

	team2 := game.Teams()[1]
	team2.Done("EC-001")
	team2.CompleteIteration()
	team2.Done("EC-002", "EC-003")
	team2.CompleteIteration()
	team2.Done("EC-004", "EC-005", "EC-006", "EC-007")
	team2.CompleteIteration()
}
