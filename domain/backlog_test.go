package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestDefaultBacklogStoriesAreNotDoneAtBeginning(t *testing.T) {
	backlog := DefaultBacklog()

	for _, story := range backlog {
		assert.False(t, story.Done)
	}
}

func TestDefaultBacklog_scores_points(t *testing.T) {
	tests := []struct {
		name           string
		storiesDone    []StoryId
		expectedPoints int
	}{
		{"score zero at beginning", []StoryId{}, 0},
		{"score when a story is done", []StoryId{"EC-001"}, 1},
		{"score when several stories are done", []StoryId{"EC-001", "EC-002", "EC-003"}, 3},
		{"do not score when story does not exist", []StoryId{"Wrong-Id"}, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backlog := DefaultBacklog()

			backlog.Done(test.storiesDone...)

			assert.Equal(t, test.expectedPoints, backlog.Score().Point)
		})
	}
}
