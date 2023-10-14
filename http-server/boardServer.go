package http_server

import (
	. "elephant_carpaccio/domain"
	"embed"
	"io/fs"
	"net/http"
	"strings"
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
	router.HandleFunc("/", s.handleBoardPage)
	router.Handle("/static/", s.staticHandler())
	router.HandleFunc("/register", s.handleRegistration)
	router.HandleFunc("/demo", s.handleDemoIndex)
	router.HandleFunc("/demo/", s.handleDemoScoring)

	s.Handler = router

	return s
}

func (s BoardServer) handleBoardPage(writer http.ResponseWriter, _ *http.Request) {
	_ = s.templateRenderer.RenderBoard(writer, s.game)
}

func (s BoardServer) handleRegistration(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		_ = s.templateRenderer.RenderRegistration(writer, s.game)
	case http.MethodPost:
		s.game.Register(request.FormValue("teamName"))
		http.Redirect(writer, request, request.URL.String(), http.StatusFound)
	}
}

func (s BoardServer) staticHandler() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.FS(staticFS)))
}

func (s BoardServer) handleDemoIndex(writer http.ResponseWriter, _ *http.Request) {
	_ = s.templateRenderer.RenderDemoIndex(writer, s.game)
}

func (s BoardServer) handleDemoScoring(writer http.ResponseWriter, request *http.Request) {
	teamName := strings.TrimPrefix(request.URL.Path, "/demo/")

	selectedTeam := s.game.FindTeamByName(teamName)

	if selectedTeam != nil {
		switch request.Method {
		case http.MethodGet:
			_ = s.templateRenderer.RenderDemoScoring(writer, selectedTeam)
		case http.MethodPost:
			var storiesDone []StoryId
			for _, story := range selectedTeam.Backlog() {
				if request.FormValue(string(story.Id)) != "" {
					storiesDone = append(storiesDone, story.Id)
				}
			}
			selectedTeam.Done(storiesDone...)
			selectedTeam.LogIterationScore()
			http.Redirect(writer, request, "/demo", http.StatusFound)
		}
	}
}
