package domain

import . "elephant_carpaccio/domain/money"

type Score struct {
	Point         int
	BusinessValue Dollar
	Risk          int
	CostOfDelay   Dollar
}

type BacklogScorer struct {
	backlog          Backlog
	currentIteration uint8
	score            *Score
}

func NewBacklogScorer(backlog Backlog, currentIteration uint8) *BacklogScorer {
	return &BacklogScorer{backlog, currentIteration, &Score{}}
}

func (t BacklogScorer) Score() Score {
	for _, story := range t.backlog {
		t.updateScoreWith(story)
	}
	return *t.score
}

func (t BacklogScorer) updateScoreWith(u UserStory) {
	t.updatePointWith(u)
	t.updateBusinessValueWith(u)
	t.updateRiskWith(u)
	t.updateCostOfDelayWith(u)
}

func (t BacklogScorer) updatePointWith(u UserStory) {
	if u.IsDone() {
		t.score.Point += u.pointEstimation
	}
}

func (t BacklogScorer) updateBusinessValueWith(u UserStory) {
	if u.IsDone() {
		t.score.BusinessValue = t.score.BusinessValue.Add(u.businessValueEstimation)
	}
}

func (t BacklogScorer) updateRiskWith(u UserStory) {
	if !u.IsDone() {
		t.score.Risk += u.riskEstimation
	}
}

func (t BacklogScorer) updateCostOfDelayWith(u UserStory) {
	nbOfWaitedIteration := t.nbOfWaitedIteration(u)
	cost := u.businessValueEstimation.Multiply(Decimal(int(nbOfWaitedIteration) * 100))
	t.score.CostOfDelay = t.score.CostOfDelay.Add(cost)
}

func (t BacklogScorer) nbOfWaitedIteration(u UserStory) uint8 {
	if u.IsDone() {
		return u.doneInIteration - u.iterationEstimation
	} else if t.currentIteration >= u.iterationEstimation {
		return t.currentIteration - u.iterationEstimation + 1
	}
	return 0
}
