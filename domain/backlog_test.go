package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain"
	. "elephant_carpaccio/domain/money"
)

func TestDefaultBacklogStoriesAreNotDoneAtBeginning(t *testing.T) {
	backlog := DefaultBacklog()

	for _, story := range backlog {
		assert.False(t, story.Done)
	}
}

func TestDefaultBacklog_add_points(t *testing.T) {
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

func TestDefaultBacklog_adds_business_value(t *testing.T) {
	tests := []struct {
		name                  string
		storiesDone           []StoryId
		expectedBusinessValue Dollar
	}{
		{"have no business value at beginning", []StoryId{}, NewDollar(Decimal(0))},
		{"add business value when a story is done", []StoryId{"EC-004"}, NewDollar(Decimal(500000))},
		{"add business value when several stories are done", []StoryId{"EC-004", "EC-005", "EC-006"}, NewDollar(Decimal(890000))},
		{"do add business value when story does not exist", []StoryId{"Wrong-Id"}, NewDollar(Decimal(0))},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backlog := DefaultBacklog()

			backlog.Done(test.storiesDone...)

			assert.Equal(t, test.expectedBusinessValue, backlog.Score().BusinessValue)
		})
	}
}

func TestDefaultBacklog_mitigates_risk(t *testing.T) {
	tests := []struct {
		name         string
		storiesDone  []StoryId
		expectedRisk int
	}{
		{"have maximum risk at beginning", []StoryId{}, 100},
		{"add business value when a story is done", []StoryId{"EC-001"}, 70},
		{"add business value when several stories are done", []StoryId{"EC-001", "EC-002", "EC-003"}, 50},
		{"do add business value when story does not exist", []StoryId{"Wrong-Id"}, 100},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backlog := DefaultBacklog()

			backlog.Done(test.storiesDone...)

			assert.Equal(t, test.expectedRisk, backlog.Score().Risk)
		})
	}
}
