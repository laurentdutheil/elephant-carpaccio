package domain

import "elephant_carpaccio/domain/money"

type StoryId string

type Score struct {
	Point         int
	BusinessValue money.Dollar
	Risk          int
}

func NewScore(point int, businessValue money.Dollar, risk int) Score {
	return Score{Point: point, BusinessValue: businessValue, Risk: risk}
}

type UserStory struct {
	Id          StoryId
	Description string
	Score       Score
	Done        bool
}

func (u UserStory) AddScoreTo(score Score) Score {
	if u.Done {
		return NewScore(u.Score.Point+score.Point, u.Score.BusinessValue.Add(score.BusinessValue), score.Risk)
	}
	return NewScore(score.Point, score.BusinessValue, u.Score.Risk+score.Risk)
}
