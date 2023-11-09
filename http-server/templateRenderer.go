package http_server

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net"

	. "elephant_carpaccio/domain"
	"elephant_carpaccio/domain/controller"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

type TemplateRenderer struct {
	template *template.Template
}

type GameBoard struct {
	Game    *Game
	BaseURL string
}

func NewTemplateRenderer() *TemplateRenderer {
	t, _ := template.ParseFS(templates, "templates/*.gohtml")
	return &TemplateRenderer{template: t}
}

func (r TemplateRenderer) RenderBoard(w io.Writer, game *Game, localIp net.IP) error {
	gameBoard := GameBoard{
		Game:    game,
		BaseURL: fmt.Sprintf("http://%v:3000", localIp.String()),
	}
	return r.template.ExecuteTemplate(w, "index.gohtml", gameBoard)
}

func (r TemplateRenderer) RenderRegistration(w io.Writer, game *Game) error {
	return r.template.ExecuteTemplate(w, "registration.gohtml", game)
}

func (r TemplateRenderer) RenderDemoIndex(w io.Writer, game *Game) error {
	return r.template.ExecuteTemplate(w, "demo-index.gohtml", game)
}

type DemoScoringModel struct {
	Team    *Team
	Order   controller.Order
	Receipt controller.Receipt
}

func (r TemplateRenderer) RenderDemoScoring(w io.Writer, team *Team) error {
	randomizer := controller.NewDecimalRandomizer(rand.Int63n)
	orderGenerator := controller.NewOrderGenerator(randomizer)
	stateCode := orderGenerator.PickStateCode()
	discountLevel := orderGenerator.PickDiscountLevel()
	order := orderGenerator.GenerateOrder(discountLevel, stateCode)
	receipt := order.Compute()

	scoringModel := DemoScoringModel{
		Team:    team,
		Order:   order,
		Receipt: receipt,
	}
	return r.template.ExecuteTemplate(w, "demo-team.gohtml", scoringModel)
}
