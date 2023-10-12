package http_server

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestScoreRender(t *testing.T) {
	game := simulateGame()

	scoreRenderer, _ := NewScoreRenderer()

	t.Run("it renders the scores of the teams", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := scoreRenderer.RenderBoard(&buf, game); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders registration form for teams", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := scoreRenderer.RenderRegistration(&buf, game); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func simulateGame() *Game {
	game := NewGame()
	game.Register("A Team")
	game.Register("The fantastic four")

	for _, team := range game.Teams() {
		team.Done("EC-001", "EC-002", "EC-003")
	}
	game.LogIteration()
	for _, team := range game.Teams() {
		team.Done("EC-004", "EC-005")
	}
	game.LogIteration()
	for _, team := range game.Teams() {
		team.Done("EC-006", "EC-007", "EC-008")
	}
	game.LogIteration()
	return game
}
