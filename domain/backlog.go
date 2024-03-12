package domain

import (
	. "elephant_carpaccio/domain/money"
)

type Backlog []UserStory

func (b Backlog) Done(currentIteration uint8, userStoryIds ...StoryId) {
	for i, story := range b {
		if contains(userStoryIds, story.Id) {
			b[i].Done(currentIteration)
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

func (b Backlog) Score(currentIteration uint8) Score {
	backlogScore := NewScore(0, NewDollar(Decimal(0)), 0, NewDollar(Decimal(0)))
	for _, story := range b {
		backlogScore = backlogScore.AddScoreOf(story, currentIteration)
	}
	return backlogScore
}

func DefaultBacklog() Backlog {
	return Backlog{
		NewUserStoryBuilder("EC-001").Description("Hello World").RiskEstimation(30).IterationEstimation(1).Build(),
		NewUserStoryBuilder("EC-002").Description("Can fill parameters").RiskEstimation(10).IterationEstimation(1).Build(),
		NewUserStoryBuilder("EC-003").Description("Compute order value without tax").BusinessValueEstimation(NewDollar(Decimal(100000))).RiskEstimation(10).IterationEstimation(1).Build(),
		NewUserStoryBuilder("EC-004").Description("Tax for UT").BusinessValueEstimation(NewDollar(Decimal(500000))).RiskEstimation(15).IterationEstimation(2).Build(),
		NewUserStoryBuilder("EC-005").Description("Tax for NV").BusinessValueEstimation(NewDollar(Decimal(260000))).RiskEstimation(5).IterationEstimation(2).Build(),
		NewUserStoryBuilder("EC-006").Description("Tax for TX").BusinessValueEstimation(NewDollar(Decimal(130000))).RiskEstimation(2).IterationEstimation(2).Build(),
		NewUserStoryBuilder("EC-007").Description("Tax for AL").BusinessValueEstimation(NewDollar(Decimal(70000))).RiskEstimation(1).IterationEstimation(2).Build(),
		NewUserStoryBuilder("EC-008").Description("Tax for CA").BusinessValueEstimation(NewDollar(Decimal(40000))).IterationEstimation(2).Build(),
		NewUserStoryBuilder("EC-009").Description("Can handle float for 'number of items' AND 'price by item'").BusinessValueEstimation(NewDollar(Decimal(100000))).RiskEstimation(10).IterationEstimation(2).Build(),
		NewUserStoryBuilder("EC-010").Description("15% Discount").BusinessValueEstimation(NewDollar(Decimal(50000))).RiskEstimation(10).IterationEstimation(3).Build(),
		NewUserStoryBuilder("EC-011").Description("10% Discount").BusinessValueEstimation(NewDollar(Decimal(26000))).RiskEstimation(2).IterationEstimation(3).Build(),
		NewUserStoryBuilder("EC-012").Description("7% Discount").BusinessValueEstimation(NewDollar(Decimal(13000))).RiskEstimation(1).IterationEstimation(3).Build(),
		NewUserStoryBuilder("EC-013").Description("5% Discount").BusinessValueEstimation(NewDollar(Decimal(7000))).IterationEstimation(3).Build(),
		NewUserStoryBuilder("EC-014").Description("3% Discount").BusinessValueEstimation(NewDollar(Decimal(4000))).IterationEstimation(3).Build(),
		NewUserStoryBuilder("EC-015").Description("Can handle rounding for result (two digits after the decimal point)").BusinessValueEstimation(NewDollar(Decimal(10000))).IterationEstimation(4).Build(),
		NewUserStoryBuilder("EC-016").Description("Prompts are clear. Display currency").BusinessValueEstimation(NewDollar(Decimal(10000))).RiskEstimation(1).IterationEstimation(4).Build(),
		NewUserStoryBuilder("EC-017").Description("Display details (order value, tax, discount").BusinessValueEstimation(NewDollar(Decimal(20000))).RiskEstimation(3).IterationEstimation(4).Build(),
		NewUserStoryBuilder("EC-018").Description("Do not have to re-launch the application for each test").IterationEstimation(5).Build(),
	}
}

type Score struct {
	Point         int
	BusinessValue Dollar
	Risk          int
	CostOfDelay   Dollar
}

func NewScore(point int, businessValue Dollar, risk int, costOfDelay Dollar) Score {
	return Score{
		Point:         point,
		BusinessValue: businessValue,
		Risk:          risk,
		CostOfDelay:   costOfDelay,
	}
}

func (s Score) AddScoreOf(u UserStory, currentIteration uint8) Score {
	return NewScore(
		s.addPoint(u),
		s.addBusinessValue(u),
		s.addRisk(u),
		s.addCostOfDelay(u, currentIteration),
	)
}

func (s Score) addPoint(u UserStory) int {
	if u.IsDone() {
		s.Point += u.pointEstimation
	}
	return s.Point
}

func (s Score) addBusinessValue(u UserStory) Dollar {
	if u.IsDone() {
		s.BusinessValue = s.BusinessValue.Add(u.businessValueEstimation)
	}
	return s.BusinessValue
}

func (s Score) addRisk(u UserStory) int {
	if !u.IsDone() {
		s.Risk += u.riskEstimation
	}
	return s.Risk
}

func (s Score) addCostOfDelay(u UserStory, currentIteration uint8) Dollar {
	var nbOfWaitedIteration int
	if u.IsDone() {
		nbOfWaitedIteration = int(u.doneInIteration - u.iterationEstimation)
	} else {
		if currentIteration >= u.iterationEstimation {
			nbOfWaitedIteration = int(currentIteration-u.iterationEstimation) + 1
		}
	}
	cost := u.businessValueEstimation.Multiply(Decimal(nbOfWaitedIteration * 100))
	return s.CostOfDelay.Add(cost)
}
