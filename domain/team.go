package domain

type Team struct {
	name            string
	ip              string
	backlog         Backlog
	iterationScores []Score
	gameNotifier    GameNotifier
}

func NewTeam(name string, gameNotifier GameNotifier) *Team {
	return &Team{name: name, backlog: DefaultBacklog(), gameNotifier: gameNotifier}
}

func (t *Team) Name() string {
	return t.name
}

func (t *Team) IP() string {
	return t.ip
}

func (t *Team) SetIp(ip string) {
	t.ip = ip
}

func (t *Team) Backlog() Backlog {
	return t.backlog
}

func (t *Team) Done(userStoryIds ...StoryId) {
	t.backlog.Done(t.currentIteration(), userStoryIds...)
}

func (t *Team) CompleteIteration() {
	currentScore := t.backlog.Score(t.currentIteration())
	t.iterationScores = append(t.iterationScores, currentScore)
	if t.gameNotifier != nil {
		t.gameNotifier.NotifyScore(t.name, currentScore)
	}
}

func (t *Team) IterationScores() []Score {
	return t.iterationScores
}

func (t *Team) currentIteration() uint8 {
	return uint8(len(t.iterationScores) + 1)
}
