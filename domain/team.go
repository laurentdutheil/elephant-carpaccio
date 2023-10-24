package domain

type Team struct {
	name            string
	backlog         Backlog
	iterationScores []int
	scoreSubject    ScoreSubject
}

func NewTeam(name string, scoreSubject ScoreSubject) *Team {
	return &Team{name: name, backlog: defaultBacklog(), scoreSubject: scoreSubject}
}

func (t *Team) Name() string {
	return t.name
}

func (t *Team) Backlog() Backlog {
	return t.backlog
}

func (t *Team) Done(userStoryIds ...StoryId) {
	for _, id := range userStoryIds {
		t.backlog.Done(id)
	}
}

func (t *Team) Score() int {
	return t.backlog.Score()
}

func (t *Team) CompleteIteration() {
	t.iterationScores = append(t.iterationScores, t.Score())
	if t.scoreSubject != nil {
		t.scoreSubject.NotifyAll(t.name, t.Score())
	}
}

func (t *Team) IterationScores() []int {
	return t.iterationScores
}
