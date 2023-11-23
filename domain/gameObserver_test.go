package domain_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestGameSubject(t *testing.T) {
	t.Run("should notify new score to all observers", func(t *testing.T) {
		mockScoreObserver := createMockedGameObserver()

		gameSubject := NewGameSubject()
		gameSubject.AddGameObserver(mockScoreObserver)

		gameSubject.NotifyScore("A Team", Score(3))

		assert.Equal(t, 1, gameSubject.NbGameObservers())
		mockScoreObserver.AssertCalled(t, "UpdateScore", "A Team", Score(3))
	})

	t.Run("should not notify score if no observer", func(t *testing.T) {
		mockScoreObserver := createMockedGameObserver()

		gameSubject := NewGameSubject()
		gameSubject.AddGameObserver(mockScoreObserver)
		gameSubject.RemoveGameObserver("ObserverId")

		gameSubject.NotifyScore("A Team", Score(3))

		assert.Equal(t, 0, gameSubject.NbGameObservers())
		mockScoreObserver.AssertNotCalled(t, "UpdateScore", mock.Anything, mock.Anything)
	})

	t.Run("should notify new registration to all observers", func(t *testing.T) {
		mockScoreObserver := createMockedGameObserver()

		gameSubject := NewGameSubject()
		gameSubject.AddGameObserver(mockScoreObserver)

		gameSubject.NotifyRegistration("A Team")

		assert.Equal(t, 1, gameSubject.NbGameObservers())
		mockScoreObserver.AssertCalled(t, "AddRegistration", "A Team")
	})

	t.Run("should not notify new registration if no observer", func(t *testing.T) {
		mockScoreObserver := createMockedGameObserver()

		gameSubject := NewGameSubject()
		gameSubject.AddGameObserver(mockScoreObserver)
		gameSubject.RemoveGameObserver("ObserverId")

		gameSubject.NotifyRegistration("A Team")

		assert.Equal(t, 0, gameSubject.NbGameObservers())
		mockScoreObserver.AssertNotCalled(t, "AddRegistration", mock.Anything)
	})
}

func createMockedGameObserver() *MockGameObserver {
	mockScoreObserver := &MockGameObserver{}
	mockScoreObserver.On("Id").Return("ObserverId")
	mockScoreObserver.On("UpdateScore", mock.Anything, mock.Anything)
	mockScoreObserver.On("AddRegistration", mock.Anything)
	return mockScoreObserver
}

type MockGameObserver struct {
	mock.Mock
}

func (m *MockGameObserver) Id() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockGameObserver) UpdateScore(teamName string, newIterationScore Score) {
	m.Called(teamName, newIterationScore)
}

func (m *MockGameObserver) AddRegistration(teamName string) {
	m.Called(teamName)
}
