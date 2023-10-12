package http_server

import (
	. "elephant_carpaccio/domain"
	"embed"
	"io/fs"
	"net/http"
)

var (
	//go:embed "static/*"
	static      embed.FS
	staticFS, _ = fs.Sub(static, "static")
)

type BoardServer struct {
	templateRenderer Renderer
	http.Handler
	game *Game
}

func NewBoardServer(renderer Renderer, game *Game) *BoardServer {
	s := &BoardServer{templateRenderer: renderer, game: game}

	router := http.NewServeMux()
	router.HandleFunc("/", s.handleRoot)
	router.HandleFunc("/register", s.handleRegistration)
	router.Handle("/static/", s.staticHandler())

	s.Handler = router

	return s
}

func (s BoardServer) handleRoot(writer http.ResponseWriter, request *http.Request) {
	_ = s.templateRenderer.RenderBoard(writer, s.game)
}

func (s BoardServer) handleRegistration(writer http.ResponseWriter, request *http.Request) {
	_ = s.templateRenderer.RenderRegistration(writer, s.game)
}

func (s BoardServer) staticHandler() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.FS(staticFS)))
}
