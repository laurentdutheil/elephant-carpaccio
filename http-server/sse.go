package http_server

import (
	"bytes"
	"elephant_carpaccio/domain"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ScoreEvent struct {
	TeamName string       `json:"teamName"`
	NewScore domain.Score `json:"newScore"`
}

type SseScoreObserver struct {
	id           string
	scoreChannel chan ScoreEvent
}

func NewSseScoreObserver() *SseScoreObserver {
	id := strconv.FormatInt(time.Now().Unix(), 10)
	scoreChannel := make(chan ScoreEvent)
	o := &SseScoreObserver{id: id, scoreChannel: scoreChannel}
	return o
}

func (o SseScoreObserver) Id() string {
	return o.id
}

func (o SseScoreObserver) Update(teamName string, newScore domain.Score) {
	o.scoreChannel <- ScoreEvent{teamName, newScore}
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
