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

func TestTeamHaveADefaultBacklogAtBeginning(t *testing.T) {
	team := NewTeam("A Team")

	backlog := team.Backlog()

	assert.Contains(t, backlog, UserStory{description: "Hello World", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Can fill parameters", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Compute order value without tax", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Can handle float for 'number of items' AND 'price by item'", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Tax for UT", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Tax for NV", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Tax for TX", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Tax for AL", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Tax for CA", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "15% Discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "10% Discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "7% Discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "5% Discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "3% Discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Can handle rounding for result (two digits after the decimal point)", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Prompts are clear. Display currency", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Display details (order value, tax, discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{description: "Do not have to re-launch the application for each test", valuePoint: 1})
}

func TestAllTheTeamBacklogAreNotDoneAtBeginning(t *testing.T) {
	team := NewTeam("A Team")

	backlog := team.Backlog()

	for _, story := range backlog {
		assert.False(t, story.done)
	}
}
