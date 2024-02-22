package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestBacklogScore_should_be_zero_if_no_story_done(t *testing.T) {
	backlog := Backlog{
		{Id: "id-1", Score: Score{Point: 1}, Done: false},
		{Id: "id-2", Score: Score{Point: 1}, Done: false},
	}

	score := backlog.Score()

	assert.Equal(t, Score{Point: 0}, score)
}

func TestBacklogScore_should_be_only_score_of_story_done(t *testing.T) {
	backlog := Backlog{
		{Id: "id-1", Score: Score{Point: 1}, Done: true},
		{Id: "id-2", Score: Score{Point: 1}, Done: true},
		{Id: "id-3", Score: Score{Point: 1}, Done: false},
	}

	score := backlog.Score()

	assert.Equal(t, Score{Point: 2}, score)
}

func TestDefaultBacklogStoriesAreNotDoneAtBeginning(t *testing.T) {

	backlog := DefaultBacklog()

	for _, story := range backlog {
		assert.False(t, story.Done)
	}
}

func TestDefaultBacklogScoresZeroAtBeginning(t *testing.T) {
	backlog := DefaultBacklog()

	assert.Equal(t, Score{Point: 0}, backlog.Score())
}

func TestDefaultBacklogScoresWhenAStoryIsDone(t *testing.T) {
	backlog := DefaultBacklog()

	backlog.Done("EC-001")

	assert.Equal(t, Score{Point: 1}, backlog.Score())
}

func TestDefaultBacklogScoresWhenSeveralStoriesAreDone(t *testing.T) {
	backlog := DefaultBacklog()

	backlog.Done("EC-001", "EC-002", "EC-003")

	assert.Equal(t, Score{Point: 3}, backlog.Score())
}

func TestDefaultBacklogDoesNotScoreWhenStoryDoesNotExist(t *testing.T) {
	backlog := DefaultBacklog()

	backlog.Done("Wrong-Id")

	assert.Equal(t, Score{Point: 0}, backlog.Score())
}
