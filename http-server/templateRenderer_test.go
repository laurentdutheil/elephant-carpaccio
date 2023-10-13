package http_server

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestTemplateRender(t *testing.T) {
	game := simulateGame()

	scoreRenderer, _ := NewRenderer()

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

	t.Run("it renders list of teams for demo", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := scoreRenderer.RenderDemoIndex(&buf, game); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders backlog of a team for demo", func(t *testing.T) {
		buf := bytes.Buffer{}

		team := game.Teams()[0]
		if err := scoreRenderer.RenderDemoScoring(&buf, team); err != nil {
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
