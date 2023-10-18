package http_server

import (
	"embed"
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
	template *template.Template
}

func NewRenderer() *TemplateRenderer {
	t, _ := template.ParseFS(templates, "templates/*.gohtml")
	return &TemplateRenderer{template: t}
}

func (r TemplateRenderer) RenderBoard(w io.Writer, game *Game) error {
	return r.template.ExecuteTemplate(w, "index.gohtml", game)
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
