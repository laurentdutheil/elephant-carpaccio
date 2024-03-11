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
		assert.False(t, story.IsDone())
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

			backlog.Done(0, test.storiesDone...)

			assert.Equal(t, test.expectedPoints, backlog.Score(0).Point)
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
		{"do not add business value when story does not exist", []StoryId{"Wrong-Id"}, NewDollar(Decimal(0))},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backlog := DefaultBacklog()

			backlog.Done(0, test.storiesDone...)

			assert.Equal(t, test.expectedBusinessValue, backlog.Score(0).BusinessValue)
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
		{"remove risk when a story is done", []StoryId{"EC-001"}, 70},
		{"remove risk when several stories are done", []StoryId{"EC-001", "EC-002", "EC-003"}, 50},
		{"do not remove risk when story does not exist", []StoryId{"Wrong-Id"}, 100},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backlog := DefaultBacklog()

			backlog.Done(0, test.storiesDone...)

			assert.Equal(t, test.expectedRisk, backlog.Score(0).Risk)
		})
	}
}

func TestDefaultBacklog_add_cost_of_delay(t *testing.T) {
	tests := []struct {
		name             string
		storiesDone      []StoryId
		currentIteration uint8
		expectedCOD      Dollar
	}{
		{"have no cost of delay at beginning", []StoryId{}, 0, NewDollar(Decimal(0))},
		{"cost of delay for the last iteration is sum of business value not done on iteration estimation", []StoryId{}, 5, NewDollar(Decimal(5280000))},
		{"cost of delay for the 4th iteration is sum of business value not done on iteration estimation", []StoryId{}, 4, NewDollar(Decimal(3940000))},
		{"cost of delay for the 3rd iteration is sum of business value not done on iteration estimation", []StoryId{}, 3, NewDollar(Decimal(2600000))},
		{"cost of delay for the 2nd iteration is sum of business value not done on iteration estimation", []StoryId{}, 2, NewDollar(Decimal(1300000))},
		{"cost of delay for the 1st iteration is sum of business value not done on iteration estimation", []StoryId{}, 1, NewDollar(Decimal(100000))},
		{"no cost of delay if stories are done in iteration estimation", []StoryId{"EC-001", "EC-002", "EC-003"}, 1, NewDollar(Decimal(0))},
		{"cost of delay if stories are done after iteration estimation", []StoryId{"EC-001", "EC-002", "EC-003"}, 2, NewDollar(Decimal(1200000))},
		{"do not change the cost of delay when story does not exist", []StoryId{"Wrong-Id"}, 5, NewDollar(Decimal(5280000))},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backlog := DefaultBacklog()

			backlog.Done(test.currentIteration, test.storiesDone...)

			assert.Equal(t, test.expectedCOD, backlog.Score(test.currentIteration).CostOfDelay)
		})
	}
}
