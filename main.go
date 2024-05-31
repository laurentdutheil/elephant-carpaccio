package main

import (
	. "elephant_carpaccio/domain"
	httpserver "elephant_carpaccio/http-server"
	"elephant_carpaccio/http-server/network"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	os.Exit(realMain(net.InterfaceAddrs, http.ListenAndServe, os.Stdout))
}

func realMain(interfaceAddrsFunc network.InterfaceAddrs, listenAndServeFunc func(addr string, handler http.Handler) error, stdout io.Writer) int {
	environment := handleFlags(stdout)

	localIp, err := network.GetLocalIp(interfaceAddrsFunc)
	if err != nil {
		_, _ = fmt.Fprintln(stdout, err.Error())
		return 1
	}
	_, _ = fmt.Fprintln(stdout, "local IP: "+localIp.String())

	game := NewGame()
	server := httpserver.NewBoardServer(game, localIp)

	addr := ":3000"
	if environment == "DEV" {
		simulateGameForDev(game)
		addr = "localhost:3000"
	}
	err = listenAndServeFunc(addr, server)
	if err != nil {
		_, _ = fmt.Fprintln(stdout, "error during launch of the server: "+err.Error())
		return 1
	}

	_, _ = fmt.Fprintln(stdout, "visit http://localhost:3000")
	return 0
}

func handleFlags(stdout io.Writer) string {
	var environment string
	flag.StringVar(&environment, "env", "PRD", "environment: DEV (development), PRD (production)")
	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "env" && f.Value.String() != "DEV" {
			environment = "PRD"
		}
	})
	_, _ = fmt.Fprintln(stdout, "env: "+environment)
	return environment
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
