package http_server_test

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"net"
	"regexp"
	"testing"

	. "elephant_carpaccio/domain"
	. "elephant_carpaccio/http-server"
)

func TestTemplateRender(t *testing.T) {
	game := simulateGame()
	templateRenderer := NewTemplateRenderer()

	t.Run("it renders the top", func(t *testing.T) {
		buf := bytes.Buffer{}

		mainScrubber, _ := regexp.Compile("(?s)<main>.*</html>")
		ignoreMainAndFooter := approvals.Options().
			WithRegexScrubber(mainScrubber, "<<main and footer templates>>")

		if err := templateRenderer.RenderBoard(&buf, game, net.ParseIP("128.168.0.44")); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String(), ignoreMainAndFooter)
	})

	t.Run("it renders the footer", func(t *testing.T) {
		buf := bytes.Buffer{}

		mainScrubber, _ := regexp.Compile("(?s)<!DOCTYPE html>.*</main>")
		ignoreTopAndMain := approvals.Options().
			WithRegexScrubber(mainScrubber, "<<top and main templates>>")

		if err := templateRenderer.RenderBoard(&buf, game, net.ParseIP("128.168.0.44")); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String(), ignoreTopAndMain)
	})

	t.Run("it renders the scores of the teams", func(t *testing.T) {
		buf := bytes.Buffer{}

		topScrubber, _ := regexp.Compile("(?s)<!DOCTYPE html>.*<main>")
		footerScrubber, _ := regexp.Compile("(?s)</main>.*</html>")
		ignoreTopAndFooter := approvals.Options().
			WithRegexScrubber(topScrubber, "<<top template>>").
			WithRegexScrubber(footerScrubber, "<<footer template>>")

		if err := templateRenderer.RenderBoard(&buf, game, net.ParseIP("128.168.0.44")); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String(), ignoreTopAndFooter)
	})

	t.Run("it renders registration form for teams", func(t *testing.T) {
		buf := bytes.Buffer{}

		topScrubber, _ := regexp.Compile("(?s)<!DOCTYPE html>.*<main>")
		footerScrubber, _ := regexp.Compile("(?s)</main>.*</html>")
		ignoreTopAndFooter := approvals.Options().
			WithRegexScrubber(topScrubber, "<<top template>>").
			WithRegexScrubber(footerScrubber, "<<footer template>>")

		if err := templateRenderer.RenderRegistration(&buf, game); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String(), ignoreTopAndFooter)
	})

	t.Run("it renders list of teams for demo", func(t *testing.T) {
		buf := bytes.Buffer{}

		topScrubber, _ := regexp.Compile("(?s)<!DOCTYPE html>.*<main>")
		footerScrubber, _ := regexp.Compile("(?s)</main>.*</html>")
		ignoreTopAndFooter := approvals.Options().
			WithRegexScrubber(topScrubber, "<<top template>>").
			WithRegexScrubber(footerScrubber, "<<footer template>>")

		if err := templateRenderer.RenderDemoIndex(&buf, game); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String(), ignoreTopAndFooter)
	})

	t.Run("it renders backlog of a team for demo", func(t *testing.T) {
		buf := bytes.Buffer{}

		topScrubber, _ := regexp.Compile("(?s)<!DOCTYPE html>.*<main>")
		footerScrubber, _ := regexp.Compile("(?s)</main>.*</html>")
		ignoreTopAndFooter := approvals.Options().
			WithRegexScrubber(topScrubber, "<<top template>>").
			WithRegexScrubber(footerScrubber, "<<footer template>>")

		team := game.Teams()[0]
		if err := templateRenderer.RenderDemoScoring(&buf, team); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String(), ignoreTopAndFooter)
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
