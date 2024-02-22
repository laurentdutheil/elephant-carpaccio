package domain

type Score struct {
	Point int
}

func (s Score) Add(score Score) Score {
	return Score{s.Point + score.Point}
}
