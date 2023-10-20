package http_server

import (
	"elephant_carpaccio/http-server/network"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestBoardServer(t *testing.T) {
	localIpSeekerStub := createLocalIpSeekerStub("128.168.0.44")

	t.Run("handle board page", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)

		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, response.Body.String(), "<canvas id=\"iterationScores\"></canvas>")
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
		// test the redirection to the register page
		assert.Equal(t, "/register", response.Result().Header.Get("Location"))
		assert.Equal(t, http.StatusFound, response.Code)
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
		game.Register("A Team")

		request, _ := http.NewRequest(http.MethodGet, "/demo/A Team", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, response.Body.String(), "<caption>Witch user stories are done?</caption>")
	})

	t.Run("handle demo scoring post", func(t *testing.T) {
		game := NewGame()
		server := NewBoardServer(game, localIpSeekerStub)
		game.Register("A Team")

		data := url.Values{}
		data.Add("Done", "EC-001")
		data.Add("Done", "EC-002")
		data.Add("Done", "EC-003")
		data.Add("Done", "EC-004")
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

func createLocalIpSeekerStub(expectedLocalIp string) network.InterfaceAddrs {
	return func() ([]net.Addr, error) {
		return []net.Addr{
			&net.IPNet{IP: net.ParseIP(expectedLocalIp)},
		}, nil
	}
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
