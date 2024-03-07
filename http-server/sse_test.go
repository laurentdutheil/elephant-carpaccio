package http_server_test

import (
	"context"
	"elephant_carpaccio/domain"
	"elephant_carpaccio/http-server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSse(t *testing.T) {
	t.Run("return error 500 when SSE is not supported", func(t *testing.T) {
		game := domain.NewGame()
		router := http.NewServeMux()
		http_server.HandleSSE(router, game)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		response := &NonSseSupportedResponseWriter{}

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("should set Header correctly", func(t *testing.T) {
		game := domain.NewGame()
		router := http.NewServeMux()
		http_server.HandleSSE(router, game)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, "text/event-stream", response.Header().Get("Content-Type"))
		assert.Equal(t, "no-cache", response.Header().Get("Cache-Control"))
		assert.Equal(t, "keep-alive", response.Header().Get("Connection"))
	})

	t.Run("should add a GameObserver when connection is open", func(t *testing.T) {
		game := domain.NewGame()
		router := http.NewServeMux()
		http_server.HandleSSE(router, game)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		time.AfterFunc(time.Millisecond, func() {
			assert.Equal(t, 1, game.NbGameObservers())
		})

		router.ServeHTTP(response, request)

	})

	t.Run("should remove GameObserver when connection is closed", func(t *testing.T) {
		game := domain.NewGame()
		router := http.NewServeMux()
		http_server.HandleSSE(router, game)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		time.AfterFunc(time.Millisecond, func() {
			assert.Equal(t, 1, game.NbGameObservers())
		})

		router.ServeHTTP(response, request)

		assert.Equal(t, 0, game.NbGameObservers())
	})

	t.Run("should send score event when an iteration is completed", func(t *testing.T) {
		game := domain.NewGame()
		game.Register("A Team", "")
		team := game.Teams()[0]

		router := http.NewServeMux()
		http_server.HandleSSE(router, game)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		time.AfterFunc(time.Millisecond, func() {
			team.Done("EC-004", "EC-005")
			team.CompleteIteration()
		})

		router.ServeHTTP(response, request)

		assert.Equal(t, "event: score\ndata: {\"teamName\":\"A Team\",\"newScore\":2,\"newBusinessValue\":7600.00,\"newRisk\":80}\n\n", response.Body.String())
	})

	t.Run("should send registration event when an team is registered", func(t *testing.T) {
		game := domain.NewGame()
		router := http.NewServeMux()
		http_server.HandleSSE(router, game)

		request, _ := http.NewRequest(http.MethodGet, "/sse", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		time.AfterFunc(time.Millisecond, func() {
			game.Register("A Team", "")
		})

		router.ServeHTTP(response, request)

		assert.Equal(t, "event: registration\ndata: {\"teamName\":\"A Team\"}\n\n", response.Body.String())
	})
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
