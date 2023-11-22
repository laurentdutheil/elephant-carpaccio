package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRegisterTeam(t *testing.T) {
	game := NewGame()
	game.Register("A Team")

	teams := game.Teams()

	assert.Contains(t, teams, NewTeam("A Team", game))
}

func TestDontRegisterTeamIfNameIsBlank(t *testing.T) {
	game := NewGame()
	game.Register("")

	teams := game.Teams()

	assert.NotContains(t, teams, NewTeam("", game))
	assert.Len(t, teams, 0)
}

func TestFindNoTeamWhenNoRegistration(t *testing.T) {
	game := NewGame()

	team := game.FindTeamByName("A Team")

	assert.Nil(t, team)
}

func TestFindTeamWhenItIsRegistered(t *testing.T) {
	game := NewGame()
	game.Register("A Team")

	team := game.FindTeamByName("A Team")

	assert.NotNil(t, team)
	assert.Equal(t, "A Team", team.name)
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

func TestGameAsGGameSubject(t *testing.T) {
	t.Run("should notify new score to all observers", func(t *testing.T) {
		mockScoreObserver := createMockedGameObserver()

		game := NewGame()
		game.Register("A Team")
		game.AddGameObserver(mockScoreObserver)
		team := game.FindTeamByName("A Team")
		team.Done("EC-001", "EC-002", "EC-003")
		team.CompleteIteration()

		assert.Equal(t, 1, game.NbGameObservers())
		mockScoreObserver.AssertCalled(t, "UpdateScore", "A Team", Score(3))
	})

	t.Run("should not notify score if no observer", func(t *testing.T) {
		mockScoreObserver := createMockedGameObserver()

		game := NewGame()
		game.Register("A Team")
		game.AddGameObserver(mockScoreObserver)
		game.RemoveGameObserver("ObserverId")
		team := game.FindTeamByName("A Team")
		team.Done("EC-001", "EC-002", "EC-003")
		team.CompleteIteration()

		assert.Equal(t, 0, game.NbGameObservers())
		mockScoreObserver.AssertNotCalled(t, "UpdateScore", mock.Anything, mock.Anything)
	})
	t.Run("should notify new registration to all observers", func(t *testing.T) {
		mockScoreObserver := createMockedGameObserver()

		game := NewGame()
		game.AddGameObserver(mockScoreObserver)
		game.Register("A Team")

		assert.Equal(t, 1, game.NbGameObservers())
		mockScoreObserver.AssertCalled(t, "AddRegistration", "A Team")
	})

	t.Run("should not notify new registration if no observer", func(t *testing.T) {
		mockScoreObserver := createMockedGameObserver()

		game := NewGame()
		game.AddGameObserver(mockScoreObserver)
		game.RemoveGameObserver("ObserverId")
		game.Register("A Team")

		assert.Equal(t, 0, game.NbGameObservers())
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
