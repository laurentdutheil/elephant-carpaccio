package domain

type IterationLogger map[*Team][]int

func NewIterationLogger() *IterationLogger {
	return &IterationLogger{}
}

func (l IterationLogger) LogIterationScore(team *Team) {
	l[team] = append(l[team], team.Score())
}

func (l IterationLogger) Scores(team *Team) []int {
	return l[team]
}
