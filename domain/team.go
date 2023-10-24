package domain

type Team struct {
	name            string
	backlog         Backlog
	iterationScores []int
}

func NewTeam(name string) *Team {
	return &Team{name: name, backlog: defaultBacklog()}
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
}

func (t *Team) IterationScores() []int {
	return t.iterationScores
}
