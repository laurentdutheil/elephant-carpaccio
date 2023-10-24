package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTeamHaveAName(t *testing.T) {
	team := NewTeam("A Team", nil)

	teamName := team.Name()

	assert.Equal(t, "A Team", teamName)
}

func TestTeamHaveADefaultBacklogAtBeginning(t *testing.T) {
	tests := []struct {
		id          StoryId
		description string
		valuePoint  int
	}{
		{id: "EC-001", description: "Hello World", valuePoint: 1},
		{id: "EC-002", description: "Can fill parameters", valuePoint: 1},
		{id: "EC-003", description: "Compute order value without tax", valuePoint: 1},
		{id: "EC-004", description: "Can handle float for 'number of items' AND 'price by item'", valuePoint: 1},
		{id: "EC-005", description: "Tax for UT", valuePoint: 1},
		{id: "EC-006", description: "Tax for NV", valuePoint: 1},
		{id: "EC-007", description: "Tax for TX", valuePoint: 1},
		{id: "EC-008", description: "Tax for AL", valuePoint: 1},
		{id: "EC-009", description: "Tax for CA", valuePoint: 1},
		{id: "EC-010", description: "15% Discount", valuePoint: 1},
		{id: "EC-011", description: "10% Discount", valuePoint: 1},
		{id: "EC-012", description: "7% Discount", valuePoint: 1},
		{id: "EC-013", description: "5% Discount", valuePoint: 1},
		{id: "EC-014", description: "3% Discount", valuePoint: 1},
		{id: "EC-015", description: "Can handle rounding for result (two digits after the decimal point)", valuePoint: 1},
		{id: "EC-016", description: "Prompts are clear. Display currency", valuePoint: 1},
		{id: "EC-017", description: "Display details (order value, tax, discount", valuePoint: 1},
		{id: "EC-018", description: "Do not have to re-launch the application for each test", valuePoint: 1},
	}

	team := NewTeam("A Team", nil)
	backlog := team.Backlog()

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			assert.Contains(t, backlog, UserStory{Id: test.id, Description: test.description, valuePoint: test.valuePoint})
		})
	}
}

func TestAllTheTeamBacklogAreNotDoneAtBeginning(t *testing.T) {
	team := NewTeam("A Team", nil)

	backlog := team.Backlog()

	for _, story := range backlog {
		assert.False(t, story.Done)
	}
}

func TestTeamScoresZeroAtBeginning(t *testing.T) {
	team := NewTeam("A Team", nil)

	assert.Equal(t, 0, team.Score())
}

func TestTeamScoresWhenAStoryIsDone(t *testing.T) {
	team := NewTeam("A Team", nil)

	team.Done("EC-001")

	assert.Equal(t, 1, team.Score())
}

func TestTeamScoresWhenSeveralStoriesAreDone(t *testing.T) {
	team := NewTeam("A Team", nil)

	team.Done("EC-001", "EC-002", "EC-003")

	assert.Equal(t, 3, team.Score())
}

func TestTeamDoesNotScoreWhenStoryDoesNotExist(t *testing.T) {
	team := NewTeam("A Team", nil)

	team.Done("Wrong-Id")

	assert.Equal(t, 0, team.Score())
}

func TestCompleteFirstIteration(t *testing.T) {
	team := NewTeam("A Team", nil)
	team.Done("EC-001")
	team.CompleteIteration()

	scores := team.IterationScores()

	assert.Equal(t, []int{1}, scores)
}

func TestCompleteSeveralIterations(t *testing.T) {
	team := NewTeam("A Team", nil)
	team.Done("EC-001")
	team.CompleteIteration()
	team.Done("EC-002", "EC-003")
	team.CompleteIteration()
	team.Done("EC-004", "EC-005", "EC-006")
	team.CompleteIteration()

	scores := team.IterationScores()

	assert.Equal(t, []int{1, 3, 6}, scores)
}

func TestCompleteIterationNotifyScoresListeners(t *testing.T) {
	mockScoreSubject := MockScoreSubject{}
	team := NewTeam("A Team", &mockScoreSubject)
	team.Done("EC-001")

	mockScoreSubject.On("NotifyAll", mock.Anything, mock.Anything)

	team.CompleteIteration()

	mockScoreSubject.AssertCalled(t, "NotifyAll", "A Team", 1)
}

func TestCompleteIterationDontNotifyIfThereIsNoScoreSubject(t *testing.T) {
	notInjectedMockScoreSubject := MockScoreSubject{}
	team := NewTeam("A Team", nil)
	team.Done("EC-001")

	notInjectedMockScoreSubject.On("NotifyAll", mock.Anything, mock.Anything)

	team.CompleteIteration()

	notInjectedMockScoreSubject.AssertNotCalled(t, "NotifyAll", "A Team", 1)
}

type MockScoreSubject struct {
	mock.Mock
}

func (m *MockScoreSubject) NotifyAll(teamName string, newIterationScore int) {
	m.Called(teamName, newIterationScore)
}
