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

	assert.Contains(t, backlog, UserStory{id: "EC-001", description: "Hello World", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-002", description: "Can fill parameters", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-003", description: "Compute order value without tax", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-004", description: "Can handle float for 'number of items' AND 'price by item'", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-005", description: "Tax for UT", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-006", description: "Tax for NV", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-007", description: "Tax for TX", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-008", description: "Tax for AL", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-009", description: "Tax for CA", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-010", description: "15% Discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-011", description: "10% Discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-012", description: "7% Discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-013", description: "5% Discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-014", description: "3% Discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-015", description: "Can handle rounding for result (two digits after the decimal point)", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-016", description: "Prompts are clear. Display currency", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-017", description: "Display details (order value, tax, discount", valuePoint: 1})
	assert.Contains(t, backlog, UserStory{id: "EC-018", description: "Do not have to re-launch the application for each test", valuePoint: 1})
}

func TestAllTheTeamBacklogAreNotDoneAtBeginning(t *testing.T) {
	team := NewTeam("A Team")

	backlog := team.Backlog()

	for _, story := range backlog {
		assert.False(t, story.done)
	}
}
