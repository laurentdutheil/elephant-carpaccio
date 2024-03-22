package http_server

import (
	"embed"
	"io/fs"
	"net"
	"net/http"
	"strconv"
	"strings"

	. "elephant_carpaccio/domain"
	. "elephant_carpaccio/domain/calculator"
)

var (
	//go:embed "static/*"
	static      embed.FS
	staticFS, _ = fs.Sub(static, "static")
)

type BoardServer struct {
	http.Handler
	templateRenderer *TemplateRenderer
	game             *Game
	localIp          net.IP
}

func NewBoardServer(game *Game, localIp net.IP) *BoardServer {
	s := &BoardServer{
		templateRenderer: NewTemplateRenderer(),
		game:             game,
		localIp:          localIp,
	}

	router := http.NewServeMux()
	router.HandleFunc("/", s.handleBoardPage)
	router.Handle("/static/", s.staticHandler())
	router.HandleFunc("/register", s.handleRegistration)
	router.HandleFunc("/demo", s.handleDemoIndex)
	router.HandleFunc("/demo/", s.handleDemoScoring)

	HandleSSE(router, game)
	HandleApi(router, game)

	s.Handler = router

	return s
}

func (s BoardServer) handleBoardPage(writer http.ResponseWriter, _ *http.Request) {
	_ = s.templateRenderer.RenderBoard(writer, s.game, s.localIp)
}

func (s BoardServer) handleRegistration(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		_ = s.templateRenderer.RenderRegistration(writer, s.game)
	case http.MethodPost:
		s.game.Register(request.FormValue("teamName"), "")
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
			stateInRequest := request.URL.Query().Get("state")
			state := requestState(stateInRequest)

			discountInRequest := request.URL.Query().Get("discount")
			discount := requestDiscount(discountInRequest)

			withoutDecimalsInRequest := request.URL.Query().Get("withoutDecimal")
			withoutDecimals, err := strconv.ParseBool(withoutDecimalsInRequest)
			if err != nil {
				withoutDecimals = false
			}

			orderBuilder := NewRandomOrderBuilder(NewOrderRandomizer()).
				WithDiscount(discount).
				WithState(state).
				WithoutDecimals(withoutDecimals)

			randomOrder := orderBuilder.Build()
			_ = s.templateRenderer.RenderDemoScoring(writer, selectedTeam, randomOrder)
		case http.MethodPost:
			storiesDone := s.extractStoryIdsSelected(request)
			selectedTeam.Done(storiesDone...)
			selectedTeam.CompleteIteration()
			http.Redirect(writer, request, "/demo", http.StatusFound)
		}
	}
}

func requestState(stateInRequest string) *State {
	if stateInRequest != "" {
		parsedStateCode, err := strconv.Atoi(stateInRequest)
		if err != nil {
			return nil
		}
		return StateOf(parsedStateCode)
	}
	return nil
}

func requestDiscount(discountInRequest string) *Discount {
	if discountInRequest != "" {
		parsedDiscountLevel, err := strconv.Atoi(discountInRequest)
		if err != nil {
			return nil
		}
		return DiscountOf(parsedDiscountLevel)
	}
	return nil
}

func (s BoardServer) extractStoryIdsSelected(request *http.Request) []StoryId {
	_ = request.ParseForm()
	var storiesDone []StoryId
	values := request.Form["Done"]
	for _, selectedStoryId := range values {
		storiesDone = append(storiesDone, StoryId(selectedStoryId))
	}
	return storiesDone
}
