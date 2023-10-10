package domain

import "fmt"

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

func (l IterationLogger) String() string {
	result := ""
	for team, scores := range l {
		result += team.name + ": "
		result += fmt.Sprintln(scores)
	}
	return result
}
