package http_server

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestBoardServer(t *testing.T) {

	t.Run("handle board page", func(t *testing.T) {
		mockRenderer := &MockRenderer{}
		game := NewGame()
		server := NewBoardServer(mockRenderer, game)

		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		mockRenderer.On("RenderBoard", response, game).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		mockRenderer.AssertExpectations(t)
	})

	t.Run("handle static files", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(nil, game)

		request, _ := http.NewRequest(http.MethodGet, "/static/css/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("handle registration page", func(t *testing.T) {
		mockRenderer := &MockRenderer{}
		game := NewGame()
		server := NewBoardServer(mockRenderer, game)

		request, _ := http.NewRequest(http.MethodGet, "/register", nil)
		response := httptest.NewRecorder()

		mockRenderer.On("RenderRegistration", response, game).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		mockRenderer.AssertExpectations(t)
	})

	t.Run("handle registration post", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(nil, game)

		data := url.Values{}
		data.Set("teamName", "A Team")
		request, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, "A Team", game.Teams()[0].Name())
		// test the redirection to the register page
		assert.Equal(t, "/register", response.Result().Header.Get("Location"))
		assert.Equal(t, http.StatusFound, response.Code)
	})

	t.Run("handle demo index page", func(t *testing.T) {
		mockRenderer := &MockRenderer{}
		game := NewGame()
		server := NewBoardServer(mockRenderer, game)

		request, _ := http.NewRequest(http.MethodGet, "/demo", nil)
		response := httptest.NewRecorder()

		mockRenderer.On("RenderDemoIndex", response, game).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		mockRenderer.AssertExpectations(t)
	})

	t.Run("handle demo scoring page for a team", func(t *testing.T) {
		mockRenderer := &MockRenderer{}
		game := NewGame()
		server := NewBoardServer(mockRenderer, game)
		game.Register("A Team")

		request, _ := http.NewRequest(http.MethodGet, "/demo/A Team", nil)
		response := httptest.NewRecorder()

		team := game.Teams()[0]
		mockRenderer.On("RenderDemoScoring", response, team).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		mockRenderer.AssertExpectations(t)
	})

	t.Run("handle demo scoring post", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(nil, game)
		game.Register("A Team")

		data := url.Values{}
		data.Set("EC-001", "EC-001")
		data.Set("EC-002", "EC-002")
		data.Set("EC-003", "EC-003")
		data.Set("EC-004", "EC-004")
		request, _ := http.NewRequest(http.MethodPost, "/demo/A Team", strings.NewReader(data.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, "A Team: [4]\n", game.PrintBoard())
		team := game.Teams()[0]
		assertStoriesDone(t, team, []StoryId{"EC-001", "EC-002", "EC-003", "EC-004"})
		// test the redirection to the register page
		assert.Equal(t, "/demo", response.Result().Header.Get("Location"))
		assert.Equal(t, http.StatusFound, response.Code)
	})

}

func assertStoriesDone(t *testing.T, team *Team, storyIds []StoryId) {
	var storiesDone []UserStory
	for _, story := range team.Backlog() {
		if story.Done {
			storiesDone = append(storiesDone, story)
		}
	}
	for _, story := range storiesDone {
		assert.Contains(t, storyIds, story.Id)
	}
}

type MockRenderer struct {
	mock.Mock
}

func (m *MockRenderer) RenderBoard(w io.Writer, game *Game) error {
	args := m.Called(w, game)
	return args.Error(0)
}

func (m *MockRenderer) RenderRegistration(w io.Writer, game *Game) error {
	args := m.Called(w, game)
	return args.Error(0)
}

func (m *MockRenderer) RenderDemoIndex(w io.Writer, game *Game) error {
	args := m.Called(w, game)
	return args.Error(0)
}

func (m *MockRenderer) RenderDemoScoring(w io.Writer, team *Team) error {
	args := m.Called(w, team)
	return args.Error(0)
}
