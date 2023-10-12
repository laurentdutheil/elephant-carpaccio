package http_server

import (
	. "elephant_carpaccio/domain"
	"net/http"
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

	s.Handler = router

	return s
}

func (s BoardServer) handleRoot(writer http.ResponseWriter, request *http.Request) {
	_ = s.templateRenderer.RenderBoard(writer, s.game)
}

func (s BoardServer) handleRegistration(writer http.ResponseWriter, request *http.Request) {
	_ = s.templateRenderer.RenderRegistration(writer, s.game)
}
