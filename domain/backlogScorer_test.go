package domain_test

import (
	"elephant_carpaccio/domain"
	"elephant_carpaccio/domain/money"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScorer_add_points(t *testing.T) {
	tests := []struct {
		name           string
		storiesDone    []domain.StoryId
		expectedPoints int
	}{
		{"score zero at beginning", []domain.StoryId{}, 0},
		{"score when a story is done", []domain.StoryId{"EC-001"}, 1},
		{"score when several stories are done", []domain.StoryId{"EC-001", "EC-002", "EC-003"}, 3},
		{"do not score when story does not exist", []domain.StoryId{"Wrong-Id"}, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backlog := domain.DefaultBacklog()
			backlog.Done(1, test.storiesDone...)

			scorer := domain.NewBacklogScorer(backlog, 1)

			assert.Equal(t, test.expectedPoints, scorer.Score().Point)
		})
	}
}

func TestScorer_adds_business_value(t *testing.T) {
	tests := []struct {
		name                  string
		storiesDone           []domain.StoryId
		expectedBusinessValue money.Dollar
	}{
		{"have no business value at beginning", []domain.StoryId{}, money.NewDollar(money.Decimal(0))},
		{"add business value when a story is done", []domain.StoryId{"EC-004"}, money.NewDollar(money.Decimal(500000))},
		{"add business value when several stories are done", []domain.StoryId{"EC-004", "EC-005", "EC-006"}, money.NewDollar(money.Decimal(890000))},
		{"do not add business value when story does not exist", []domain.StoryId{"Wrong-Id"}, money.NewDollar(money.Decimal(0))},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backlog := domain.DefaultBacklog()
			backlog.Done(1, test.storiesDone...)

			scorer := domain.NewBacklogScorer(backlog, 1)

			assert.Equal(t, test.expectedBusinessValue, scorer.Score().BusinessValue)
		})
	}
}

func TestScorer_mitigates_risk(t *testing.T) {
	tests := []struct {
		name         string
		storiesDone  []domain.StoryId
		expectedRisk int
	}{
		{"have maximum risk at beginning", []domain.StoryId{}, 100},
		{"remove risk when a story is done", []domain.StoryId{"EC-001"}, 70},
		{"remove risk when several stories are done", []domain.StoryId{"EC-001", "EC-002", "EC-003"}, 35},
		{"do not remove risk when story does not exist", []domain.StoryId{"Wrong-Id"}, 100},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backlog := domain.DefaultBacklog()
			backlog.Done(1, test.storiesDone...)

			scorer := domain.NewBacklogScorer(backlog, 1)

			assert.Equal(t, test.expectedRisk, scorer.Score().Risk)
		})
	}
}

func TestScorer_add_cost_of_delay(t *testing.T) {
	tests := []struct {
		name             string
		storiesDone      []domain.StoryId
		currentIteration uint8
		expectedCOD      money.Dollar
	}{
		{"have no cost of delay at beginning", []domain.StoryId{}, 0, money.NewDollar(money.Decimal(0))},
		{"cost of delay for the last iteration is sum of business value not done on iteration estimation", []domain.StoryId{}, 5, money.NewDollar(money.Decimal(5280000))},
		{"cost of delay for the 4th iteration is sum of business value not done on iteration estimation", []domain.StoryId{}, 4, money.NewDollar(money.Decimal(3940000))},
		{"cost of delay for the 3rd iteration is sum of business value not done on iteration estimation", []domain.StoryId{}, 3, money.NewDollar(money.Decimal(2600000))},
		{"cost of delay for the 2nd iteration is sum of business value not done on iteration estimation", []domain.StoryId{}, 2, money.NewDollar(money.Decimal(1300000))},
		{"cost of delay for the 1st iteration is sum of business value not done on iteration estimation", []domain.StoryId{}, 1, money.NewDollar(money.Decimal(100000))},
		{"no cost of delay if stories are done in iteration estimation", []domain.StoryId{"EC-001", "EC-002", "EC-003"}, 1, money.NewDollar(money.Decimal(0))},
		{"cost of delay if stories are done after iteration estimation", []domain.StoryId{"EC-001", "EC-002", "EC-003"}, 2, money.NewDollar(money.Decimal(1200000))},
		{"do not change the cost of delay when story does not exist", []domain.StoryId{"Wrong-Id"}, 5, money.NewDollar(money.Decimal(5280000))},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backlog := domain.DefaultBacklog()
			backlog.Done(test.currentIteration, test.storiesDone...)

			scorer := domain.NewBacklogScorer(backlog, test.currentIteration)

			assert.Equal(t, test.expectedCOD, scorer.Score().CostOfDelay)
		})
	}
}
