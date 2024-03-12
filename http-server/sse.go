package http_server

import (
	"bytes"
	"elephant_carpaccio/domain"
	"elephant_carpaccio/domain/money"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HandleSSE(router *http.ServeMux, game *domain.Game) {
	router.Handle("/sse", handleSse(game))
}

type ScoreEvent struct {
	TeamName         string        `json:"teamName"`
	NewScore         int           `json:"newScore"`
	NewBusinessValue money.Decimal `json:"newBusinessValue"`
	NewRisk          int           `json:"newRisk"`
	NewCostOfDelay   money.Decimal `json:"newCostOfDelay"`
}

type RegistrationEvent struct {
	TeamName string `json:"teamName"`
}

func handleSse(game *domain.Game) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		flusher, ok := writer.(http.Flusher)
		if !ok {
			http.Error(writer, "SSE not supported", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "text/event-stream")
		writer.Header().Set("Cache-Control", "no-cache")
		writer.Header().Set("Connection", "keep-alive")

		gameObserver := newSseGameObserver()
		game.AddGameObserver(gameObserver)

		for {
			select {
			case <-request.Context().Done():
				close(gameObserver.scoreChannel)
				game.RemoveGameObserver(gameObserver.Id())
				return
			case scoreEvent := <-gameObserver.scoreChannel:
				_, _ = fmt.Fprint(writer, formatSseEvent("score", scoreEvent))
				flusher.Flush()
			case registrationEvent := <-gameObserver.registrationChannel:
				_, _ = fmt.Fprint(writer, formatSseEvent("registration", registrationEvent))
				flusher.Flush()
			}
		}
	})
}

func formatSseEvent(event string, data any) string {
	buff := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buff)
	_ = encoder.Encode(data)

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("event: %s\n", event))
	sb.WriteString(fmt.Sprintf("data: %v\n", buff.String()))

	return sb.String()
}

type sseGameObserver struct {
	id                  string
	scoreChannel        chan ScoreEvent
	registrationChannel chan RegistrationEvent
}

func newSseGameObserver() *sseGameObserver {
	return &sseGameObserver{
		id:                  strconv.FormatInt(time.Now().Unix(), 10),
		scoreChannel:        make(chan ScoreEvent),
		registrationChannel: make(chan RegistrationEvent),
	}
}

func (o sseGameObserver) Id() string {
	return o.id
}

func (o sseGameObserver) UpdateScore(teamName string, newScore domain.Score) {
	o.scoreChannel <- ScoreEvent{
		TeamName:         teamName,
		NewScore:         newScore.Point,
		NewBusinessValue: newScore.BusinessValue.AmountInCents(),
		NewRisk:          newScore.Risk,
		NewCostOfDelay:   newScore.CostOfDelay.AmountInCents(),
	}
}

func (o sseGameObserver) AddRegistration(teamName string) {
	o.registrationChannel <- RegistrationEvent{TeamName: teamName}
}
