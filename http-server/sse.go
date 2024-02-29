package http_server

import (
	"bytes"
	"elephant_carpaccio/domain"
	"elephant_carpaccio/domain/money"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ScoreEvent struct {
	TeamName         string        `json:"teamName"`
	NewScore         int           `json:"newScore"`
	NewBusinessValue money.Decimal `json:"newBusinessValue"`
}

type RegistrationEvent struct {
	TeamName string `json:"teamName"`
}

type SseGameObserver struct {
	id                  string
	scoreChannel        chan ScoreEvent
	registrationChannel chan RegistrationEvent
}

func NewSseGameObserver() *SseGameObserver {
	id := strconv.FormatInt(time.Now().Unix(), 10)
	scoreChannel := make(chan ScoreEvent)
	registrationChannel := make(chan RegistrationEvent)
	o := &SseGameObserver{id: id, scoreChannel: scoreChannel, registrationChannel: registrationChannel}
	return o
}

func (o SseGameObserver) Id() string {
	return o.id
}

func (o SseGameObserver) UpdateScore(teamName string, newScore domain.Score) {
	o.scoreChannel <- ScoreEvent{teamName, newScore.Point, newScore.BusinessValue.AmountInCents()}
}

func (o SseGameObserver) AddRegistration(teamName string) {
	o.registrationChannel <- RegistrationEvent{teamName}
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
