package http_server_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"testing"
	"time"

	. "elephant_carpaccio/domain"
	. "elephant_carpaccio/http-server"
)

func TestBoardServer(t *testing.T) {
	localIpSeekerStub := net.ParseIP("128.168.0.44")

	t.Run("handle board page", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, response.Body.String(), "<canvas id=\"iterationScores\"></canvas>")
		assert.Contains(t, response.Body.String(), "128.168.0.44:3000/demo")
	})

	t.Run("handle static files", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/static/css/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("handle registration page", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/register", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, response.Body.String(), "<th>Registered Teams</th>")
	})

	t.Run("handle registration post", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)

		data := url.Values{}
		data.Set("teamName", "A Team")
		request, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, "A Team", game.Teams()[0].Name())
		assertRedirection(t, response, "/register")
	})

	t.Run("handle demo index page", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/demo", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, response.Body.String(), "<caption>Choose a Team for a demo</caption>")
	})

	t.Run("handle demo scoring page for a team", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)
		game.Register("A Team", "")

		request, _ := http.NewRequest(http.MethodGet, "/demo/A Team", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, response.Body.String(), "<caption>Witch user stories are done?</caption>")
	})

	t.Run("handle demo scoring page for a team with a order example with fixed state", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)
		game.Register("A Team", "")

		request, _ := http.NewRequest(http.MethodGet, "/demo/A Team?state=1", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, response.Body.String(), "Tax (8.00%)")
	})

	t.Run("handle demo scoring page for a team with a order example with wrong state", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)
		game.Register("A Team", "")
		tests := []struct {
			request string
		}{
			{"/demo/A Team?state=abc"},
			{"/demo/A Team?state=123"},
		}
		for _, test := range tests {
			t.Run(test.request, func(t *testing.T) {
				request, _ := http.NewRequest(http.MethodGet, test.request, nil)
				response := httptest.NewRecorder()

				server.ServeHTTP(response, request)

				assert.Equal(t, http.StatusOK, response.Code)
				assert.Contains(t, response.Body.String(), "<caption>Witch user stories are done?</caption>")
				assert.Regexp(t, regexp.MustCompile("Tax \\(\\d+.\\d{2}%\\)"), response.Body.String())
			})
		}
	})

	t.Run("handle demo scoring page for a team with a order with fixed discount", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)
		game.Register("A Team", "")

		request, _ := http.NewRequest(http.MethodGet, "/demo/A Team?discount=3", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, response.Body.String(), "Discount (7.00%)")
	})

	t.Run("handle demo scoring page for a team with a order example with wrong discount", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)
		game.Register("A Team", "")
		tests := []struct {
			request string
		}{
			{"/demo/A Team?discount=abc"},
			{"/demo/A Team?discount=123"},
		}
		for _, test := range tests {
			t.Run(test.request, func(t *testing.T) {
				request, _ := http.NewRequest(http.MethodGet, test.request, nil)
				response := httptest.NewRecorder()

				server.ServeHTTP(response, request)

				assert.Equal(t, http.StatusOK, response.Code)
				assert.Contains(t, response.Body.String(), "<caption>Witch user stories are done?</caption>")
				assert.Regexp(t, regexp.MustCompile("Discount \\(\\d+.\\d{2}%\\)"), response.Body.String())
			})
		}
	})

	t.Run("handle demo scoring post", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)
		game.Register("A Team", "")
		team := game.Teams()[0]

		data := url.Values{}
		data.Add("Done", "EC-001")
		data.Add("Done", "EC-002")
		data.Add("Done", "EC-003")
		data.Add("Done", "EC-004")
		request, _ := http.NewRequest(http.MethodPost, "/demo/"+team.Name(), strings.NewReader(data.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Len(t, team.IterationScores(), 1)
		assertStoriesDone(t, team.Backlog(), []StoryId{"EC-001", "EC-002", "EC-003", "EC-004"})
		assertRedirection(t, response, "/demo")
	})
}

func TestSse(t *testing.T) {
	localIpSeekerStub := net.ParseIP("128.168.0.44")

	t.Run("return error 500 when SSE is not supported", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		response := &NonSseSupportedResponseWriter{}

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("should set Header correctly", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, "text/event-stream", response.Header().Get("Content-Type"))
		assert.Equal(t, "no-cache", response.Header().Get("Cache-Control"))
		assert.Equal(t, "keep-alive", response.Header().Get("Connection"))
	})

	t.Run("should add a GameObserver when connection is open", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		time.AfterFunc(time.Millisecond, func() {
			assert.Equal(t, 1, game.NbGameObservers())
		})

		server.ServeHTTP(response, request)

	})

	t.Run("should remove GameObserver when connection is closed", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, 0, game.NbGameObservers())
	})

	t.Run("should send score event when an iteration is completed", func(t *testing.T) {
		game := NewGame()
		game.Register("A Team", "")
		team := game.Teams()[0]

		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		time.AfterFunc(time.Millisecond, func() {
			team.Done("EC-004", "EC-005")
			team.CompleteIteration()
		})

		server.ServeHTTP(response, request)

		assert.Equal(t, "event: score\ndata: {\"teamName\":\"A Team\",\"newScore\":2,\"newBusinessValue\":7600.00,\"newRisk\":80}\n\n", response.Body.String())
	})

	t.Run("should send registration event when an team is registered", func(t *testing.T) {
		game := NewGame()

		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		time.AfterFunc(time.Millisecond, func() {
			game.Register("A Team", "")
		})

		server.ServeHTTP(response, request)

		assert.Equal(t, "event: registration\ndata: {\"teamName\":\"A Team\"}\n\n", response.Body.String())
	})
}

func assertRedirection(t *testing.T, response *httptest.ResponseRecorder, expectedUrl string) {
	assert.Equal(t, expectedUrl, response.Result().Header.Get("Location"))
	assert.Equal(t, http.StatusFound, response.Code)
}

func assertStoriesDone(t *testing.T, backlog Backlog, storyIds []StoryId) {
	var storiesDone []UserStory
	for _, story := range backlog {
		if story.Done {
			storiesDone = append(storiesDone, story)
		}
	}
	for _, story := range storiesDone {
		assert.Contains(t, storyIds, story.Id)
	}
}

// A non SSE supported ResponseWriter doesn't implement http.Flusher
type NonSseSupportedResponseWriter struct {
	Code int
}

func (n *NonSseSupportedResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (n *NonSseSupportedResponseWriter) Write(_ []byte) (int, error) {
	return 0, nil
}

func (n *NonSseSupportedResponseWriter) WriteHeader(statusCode int) {
	n.Code = statusCode
}
