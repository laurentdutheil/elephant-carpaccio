package domain

type ScoreSubject interface {
	NotifyAll(teamName string, newIterationScore Score)
}
