package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogFirstIteration(t *testing.T) {
	iterationLogger := NewIterationLogger()
	team := NewTeam("A Team")
	team.Done("EC-001")
	iterationLogger.LogIterationScore(team)

	scores := iterationLogger.Scores(team)

	assert.Equal(t, []int{1}, scores)
}

func TestLogSeveralIterations(t *testing.T) {
	iterationLogger := NewIterationLogger()
	team := NewTeam("A Team")
	team.Done("EC-001")
	iterationLogger.LogIterationScore(team)
	team.Done("EC-002")
	team.Done("EC-003")
	iterationLogger.LogIterationScore(team)
	team.Done("EC-004")
	team.Done("EC-005")
	team.Done("EC-006")
	iterationLogger.LogIterationScore(team)

	scores := iterationLogger.Scores(team)

	assert.Equal(t, []int{1, 3, 6}, scores)
}
