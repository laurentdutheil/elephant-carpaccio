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

type TemplateRenderer struct {
	template *template.Template
}

func NewRenderer() (*TemplateRenderer, error) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &TemplateRenderer{template: t}, nil
}

func (r TemplateRenderer) RenderBoard(w io.Writer, game *Game) error {
	return r.template.ExecuteTemplate(w, "index.gohtml", game)
}

func (r TemplateRenderer) RenderRegistration(w io.Writer, game *Game) error {
	return r.template.ExecuteTemplate(w, "registration.gohtml", game)
}
