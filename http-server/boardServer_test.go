package http_server

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestHandleRoot(t *testing.T) {
	mockRenderer := &MockRenderer{}
	game := NewGame()
	server := NewBoardServer(mockRenderer, game)
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	mockRenderer.On("RenderBoard", response, game).Return(nil)

	server.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusOK)
	mockRenderer.AssertExpectations(t)
}

func TestHandleRegistration(t *testing.T) {
	mockRenderer := &MockRenderer{}
	game := NewGame()
	server := NewBoardServer(mockRenderer, game)
	request, _ := http.NewRequest(http.MethodGet, "/register", nil)
	response := httptest.NewRecorder()

	mockRenderer.On("RenderRegistration", response, game).Return(nil)

	server.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusOK)
	mockRenderer.AssertExpectations(t)
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
