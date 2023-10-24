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

type MockScoreObserver struct {
	mock.Mock
}

func (m *MockScoreObserver) Id() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockScoreObserver) Update(teamName string, newIterationScore Score) {
	m.Called(teamName, newIterationScore)
}

func TestGameAsScoreSubject(t *testing.T) {
	t.Run("should notify all observers", func(t *testing.T) {
		mockScoreObserver := &MockScoreObserver{}
		mockScoreObserver.On("Id").Return("ObserverId")
		mockScoreObserver.On("Update", mock.Anything, mock.Anything)

		game := NewGame()
		game.AddScoreObserver(mockScoreObserver)
		game.NotifyAll("A Team", Score(3))

		mockScoreObserver.AssertCalled(t, "Update", "A Team", Score(3))
	})

	t.Run("should not notify if no observer", func(t *testing.T) {
		mockScoreObserver := &MockScoreObserver{}
		mockScoreObserver.On("Id").Return("ObserverId")
		mockScoreObserver.On("Update", mock.Anything, mock.Anything)

		game := NewGame()
		game.AddScoreObserver(mockScoreObserver)
		game.RemoveScoreObserver("ObserverId")
		game.NotifyAll("A Team", Score(3))

		mockScoreObserver.AssertNotCalled(t, "Update", mock.Anything, mock.Anything)
	})
}
