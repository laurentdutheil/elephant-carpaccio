package http_server_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	. "elephant_carpaccio/domain"
	. "elephant_carpaccio/http-server"
)

func TestApiRouter(t *testing.T) {
	t.Run("should register team on a PUT", func(t *testing.T) {
		game := NewGame()
		router := http.NewServeMux()
		HandleApi(router, game)

		jsonBodyRequest := []byte(`{"teamName":"A Team", "ip":"128.168.0.44"}`)
		request, _ := http.NewRequest(http.MethodPut, "/api/register", bytes.NewBuffer(jsonBodyRequest))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
		teams := game.Teams()
		assert.Len(t, teams, 1)
		assert.Contains(t, teams, NewTeam("A Team", "128.168.0.44", game))
	})

	t.Run("should update registered team on a PUT", func(t *testing.T) {
		game := NewGame()
		game.Register("A Team", "128.168.0.44")
		router := http.NewServeMux()
		HandleApi(router, game)

		jsonBodyRequest := []byte(`{"teamName":"A Team", "ip":"128.168.0.55"}`)
		request, _ := http.NewRequest(http.MethodPut, "/api/register", bytes.NewBuffer(jsonBodyRequest))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		teams := game.Teams()
		assert.Len(t, teams, 1)
		assert.Contains(t, teams, NewTeam("A Team", "128.168.0.55", game))
	})

	t.Run("should not register team on a POST", func(t *testing.T) {
		game := NewGame()
		router := http.NewServeMux()
		HandleApi(router, game)

		jsonBodyRequest := []byte(`{"teamName":"A Team", "ip":"128.168.0.55"}`)
		request, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(jsonBodyRequest))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "please use PUT method to register/update your team", response.Body.String())
		teams := game.Teams()
		assert.Len(t, teams, 0)
	})

	t.Run("should not register team on bad json request", func(t *testing.T) {
		tests := []struct {
			name               string
			badJsonRequestBody []byte
		}{
			{"bad json", []byte(`{"teamName":A Team?`)},
			{"bad field for teamName", []byte(`{"team":"A Team", "ip":"128.168.0.55"}`)},
			{"bad field for ip", []byte(`{"teamName":"A Team", "ipAddress":"128.168.0.55"}`)},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				game := NewGame()
				router := http.NewServeMux()
				HandleApi(router, game)

				request, _ := http.NewRequest(http.MethodPut, "/api/register", bytes.NewBuffer(test.badJsonRequestBody))
				request.Header.Set("Content-Type", "application/json; charset=UTF-8")
				response := httptest.NewRecorder()

				router.ServeHTTP(response, request)

				assert.Equal(t, http.StatusBadRequest, response.Code)
				assert.Equal(t, "the body of your request don't respect {\"teamName\":\"<your team name>\", \"ip\":\"<your api address>\"}", response.Body.String())
				teams := game.Teams()
				assert.Len(t, teams, 0)

			})
		}

	})
}

func TestError(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TODO: test cases"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
