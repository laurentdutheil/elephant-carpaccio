package domain

type Team struct {
	name            string
	ip              string
	backlog         Backlog
	iterationScores []Score
	scoreSubject    GameNotifier
}

func NewTeam(name string, ip string, scoreSubject GameNotifier) *Team {
	return &Team{name: name, ip: ip, backlog: defaultBacklog(), scoreSubject: scoreSubject}
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
	for _, id := range userStoryIds {
		t.backlog.Done(id)
	}
}

func (t *Team) Score() Score {
	return t.backlog.Score()
}

func (t *Team) CompleteIteration() {
	t.iterationScores = append(t.iterationScores, t.Score())
	if t.scoreSubject != nil {
		t.scoreSubject.NotifyScore(t.name, t.Score())
	}
}

func (t *Team) IterationScores() []Score {
	return t.iterationScores
}
