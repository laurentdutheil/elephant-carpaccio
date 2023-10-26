package domain

type ScoreSubject interface {
	AddScoreObserver(observer ScoreObserver)
	RemoveScoreObserver(id string)
	NotifyAll(teamName string, newScore Score)
}

type ScoreObserver interface {
	Id() string
	Update(teamName string, newScore Score)
}
