package domain

import "elephant_carpaccio/domain/money"

type StoryId string

type Score struct {
	Point         int
	BusinessValue money.Dollar
}

func NewScore(point int, businessValue money.Dollar) Score {
	return Score{Point: point, BusinessValue: businessValue}
}

type UserStory struct {
	Id          StoryId
	Description string
	Score       Score
	Done        bool
}

func (u UserStory) AddScoreTo(score Score) Score {
	if u.Done {
		return NewScore(u.Score.Point+score.Point, u.Score.BusinessValue.Add(score.BusinessValue))
	}
	return NewScore(score.Point, score.BusinessValue)
}
