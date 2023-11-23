package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestRegisterTeam(t *testing.T) {
	game := NewGame()
	game.Register("A Team", "128.168.0.44")

	teams := game.Teams()

	assert.Contains(t, teams, NewTeam("A Team", "128.168.0.44", game))
}

func TestDontRegisterTeamIfNameIsBlank(t *testing.T) {
	game := NewGame()
	game.Register("", "128.168.0.44")

	teams := game.Teams()

	assert.NotContains(t, teams, NewTeam("", "128.168.0.44", game))
	assert.Len(t, teams, 0)
}

func TestFindNoTeamWhenNoRegistration(t *testing.T) {
	game := NewGame()

	team := game.FindTeamByName("A Team")

	assert.Nil(t, team)
}

func TestFindTeamWhenItIsRegistered(t *testing.T) {
	game := NewGame()
	game.Register("A Team", "")

	team := game.FindTeamByName("A Team")

	assert.NotNil(t, team)
	assert.Equal(t, "A Team", team.Name())
}

func TestRegistrationShouldNotifyAllObserver(t *testing.T) {
	mockScoreObserver := createMockedGameObserver()

	game := NewGame()
	game.AddGameObserver(mockScoreObserver)
	game.Register("A Team", "")

	mockScoreObserver.AssertCalled(t, "AddRegistration", "A Team")
}
