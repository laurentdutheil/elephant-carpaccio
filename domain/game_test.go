package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterTeam(t *testing.T) {
	game := NewGame()
	game.Register("A Team")

	teams := game.Teams()

	assert.Contains(t, teams, NewTeam("A Team"))
}

func TestDontRegisterTeamIfNameIsBlank(t *testing.T) {
	game := NewGame()
	game.Register("")

	teams := game.Teams()

	assert.NotContains(t, teams, NewTeam(""))
	assert.Len(t, teams, 0)
}

func TestLogIterationsScoreForAllTeam(t *testing.T) {
	game := NewGame()
	game.Register("A Team")
	game.Register("The fantastic four")

	for _, team := range game.Teams() {
		team.Done("EC-001", "EC-002", "EC-003")
	}
	game.LogIteration()
	for _, team := range game.Teams() {
		team.Done("EC-004", "EC-005")
	}
	game.LogIteration()
	for _, team := range game.Teams() {
		team.Done("EC-006", "EC-007", "EC-008")
	}
	game.LogIteration()

	printedBoard := game.PrintBoard()
	assert.Equal(t, "A Team: [3 5 8]\nThe fantastic four: [3 5 8]\n", printedBoard)
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
