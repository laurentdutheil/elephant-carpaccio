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

func (s Score) AddScoreOf(u UserStory) Score {
	if u.Done {
		return NewScore(u.Score.Point+s.Point, u.Score.BusinessValue.Add(s.BusinessValue), s.Risk)
	}
	return NewScore(s.Point, s.BusinessValue, u.Score.Risk+s.Risk)
}
