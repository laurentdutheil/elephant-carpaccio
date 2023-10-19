package http_server

import (
	"elephant_carpaccio/http-server/network"
	"embed"
	"fmt"
	"html/template"
	"io"

	. "elephant_carpaccio/domain"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

type Renderer interface {
	RenderBoard(w io.Writer, game *Game) error
	RenderRegistration(w io.Writer, game *Game) error
	RenderDemoIndex(w io.Writer, game *Game) error
	RenderDemoScoring(w io.Writer, team *Team) error
}

type TemplateRenderer struct {
	template       *template.Template
	interfaceAddrs network.InterfaceAddrs
}

type GameBoard struct {
	Game    *Game
	BaseURL string
}

func NewRenderer(interfaceAddr network.InterfaceAddrs) *TemplateRenderer {
	t, _ := template.ParseFS(templates, "templates/*.gohtml")
	return &TemplateRenderer{template: t, interfaceAddrs: interfaceAddr}
}

func (r TemplateRenderer) RenderBoard(w io.Writer, game *Game) error {
	localIp, _ := network.GetLocalIp(r.interfaceAddrs)
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

func (r TemplateRenderer) RenderDemoScoring(w io.Writer, team *Team) error {
	return r.template.ExecuteTemplate(w, "demo-team.gohtml", team)
}
