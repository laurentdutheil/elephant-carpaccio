package domain_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestTeamHaveAName(t *testing.T) {
	team := NewTeam("A Team", "", nil)

	teamName := team.Name()

	assert.Equal(t, "A Team", teamName)
}

func TestTeamHaveAnIP(t *testing.T) {
	team := NewTeam("A Team", "128.168.0.44", nil)

	ip := team.IP()

	assert.Equal(t, "128.168.0.44", ip)
}

func TestTeamHaveADefaultBacklogAtBeginning(t *testing.T) {
	team := NewTeam("A Team", "", nil)
	backlog := team.Backlog()

	assert.Equal(t, DefaultBacklog(), backlog)
}

func TestCompleteFirstIteration(t *testing.T) {
	team := NewTeam("A Team", "", nil)
	team.Done("EC-001")
	team.CompleteIteration()

	assert.Len(t, team.IterationScores(), 1)
}

func TestCompleteSeveralIterations(t *testing.T) {
	team := NewTeam("A Team", "", nil)
	team.Done("EC-001")
	team.CompleteIteration()
	team.Done("EC-002", "EC-003")
	team.CompleteIteration()
	team.Done("EC-004", "EC-005", "EC-006")
	team.CompleteIteration()

	assert.Len(t, team.IterationScores(), 3)
}

func TestCompleteIterationNotifyScoresListeners(t *testing.T) {
	mockScoreSubject := MockGameSubject{}
	team := NewTeam("A Team", "", &mockScoreSubject)
	team.Done("EC-001")

	mockScoreSubject.On("NotifyScore", mock.Anything, mock.Anything)

	team.CompleteIteration()

	mockScoreSubject.AssertCalled(t, "NotifyScore", "A Team", mock.AnythingOfType("Score"))
}

type MockGameSubject struct {
	mock.Mock
}

func (m *MockGameSubject) AddGameObserver(observer GameObserver) {
	m.Called(observer)
}

func (m *MockGameSubject) RemoveGameObserver(id string) {
	m.Called(id)
}

func (m *MockGameSubject) NotifyScore(teamName string, newIterationScore Score) {
	m.Called(teamName, newIterationScore)
}

func (m *MockGameSubject) NotifyRegistration(teamName string) {
	m.Called(teamName)
}
