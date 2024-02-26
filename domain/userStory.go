package domain

type StoryId string

type Score struct {
	Point int
}

type UserStory struct {
	Id          StoryId
	Description string
	Score       Score
	Done        bool
}

func (u UserStory) AddScoreTo(score Score) Score {
	if u.Done {
		return Score{u.Score.Point + score.Point}
	}
	return Score{score.Point}
}
