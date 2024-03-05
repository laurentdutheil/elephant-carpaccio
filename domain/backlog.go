package domain

import . "elephant_carpaccio/domain/money"

type Backlog []UserStory

func (b Backlog) Done(userStoryIds ...StoryId) {
	for i, story := range b {
		if contains(userStoryIds, story.Id) {
			b[i].Done = true
		}
	}
}

func contains(s []StoryId, e StoryId) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (b Backlog) Score() Score {
	backlogScore := NewScore(0, NewDollar(Decimal(0)), 0)
	for _, story := range b {
		backlogScore = story.AddScoreTo(backlogScore)
	}
	return backlogScore
}

func DefaultBacklog() Backlog {
	return Backlog{
		{Id: "EC-001", Description: "Hello World", Score: NewScore(1, NewDollar(Decimal(0)), 30)},
		{Id: "EC-002", Description: "Can fill parameters", Score: NewScore(1, NewDollar(Decimal(0)), 10)},
		{Id: "EC-003", Description: "Compute order value without tax", Score: NewScore(1, NewDollar(Decimal(100000)), 10)},
		{Id: "EC-004", Description: "Tax for UT", Score: NewScore(1, NewDollar(Decimal(500000)), 15)},
		{Id: "EC-005", Description: "Tax for NV", Score: NewScore(1, NewDollar(Decimal(260000)), 5)},
		{Id: "EC-006", Description: "Tax for TX", Score: NewScore(1, NewDollar(Decimal(130000)), 2)},
		{Id: "EC-007", Description: "Tax for AL", Score: NewScore(1, NewDollar(Decimal(70000)), 1)},
		{Id: "EC-008", Description: "Tax for CA", Score: NewScore(1, NewDollar(Decimal(40000)), 0)},
		{Id: "EC-009", Description: "Can handle float for 'number of items' AND 'price by item'", Score: NewScore(1, NewDollar(Decimal(100000)), 10)},
		{Id: "EC-010", Description: "15% Discount", Score: NewScore(1, NewDollar(Decimal(50000)), 10)},
		{Id: "EC-011", Description: "10% Discount", Score: NewScore(1, NewDollar(Decimal(26000)), 2)},
		{Id: "EC-012", Description: "7% Discount", Score: NewScore(1, NewDollar(Decimal(13000)), 1)},
		{Id: "EC-013", Description: "5% Discount", Score: NewScore(1, NewDollar(Decimal(7000)), 0)},
		{Id: "EC-014", Description: "3% Discount", Score: NewScore(1, NewDollar(Decimal(4000)), 0)},
		{Id: "EC-015", Description: "Can handle rounding for result (two digits after the decimal point)", Score: NewScore(1, NewDollar(Decimal(10000)), 0)},
		{Id: "EC-016", Description: "Prompts are clear. Display currency", Score: NewScore(1, NewDollar(Decimal(10000)), 1)},
		{Id: "EC-017", Description: "Display details (order value, tax, discount", Score: NewScore(1, NewDollar(Decimal(20000)), 3)},
		{Id: "EC-018", Description: "Do not have to re-launch the application for each test", Score: NewScore(1, NewDollar(Decimal(0)), 0)},
	}
}
