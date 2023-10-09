package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTeamHaveAName(t *testing.T) {
	team := NewTeam("A Team")

	teamName := team.Name()

	assert.Equal(t, "A Team", teamName)
}
