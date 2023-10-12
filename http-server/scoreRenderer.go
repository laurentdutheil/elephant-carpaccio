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

type ScoreRenderer struct {
	template *template.Template
}

func NewScoreRenderer() (*ScoreRenderer, error) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &ScoreRenderer{template: t}, nil
}

func (r ScoreRenderer) RenderBoard(w io.Writer, game *Game) error {
	return r.template.ExecuteTemplate(w, "index.gohtml", game)
}

func (r ScoreRenderer) RenderRegistration(w io.Writer, game *Game) error {
	return r.template.ExecuteTemplate(w, "registration.gohtml", game)
}
