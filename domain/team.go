package domain

type Team struct {
	name            string
	ip              string
	backlog         Backlog
	iterationScores []Score
	scoreSubject    GameNotifier
}

func NewTeam(name string, ip string, scoreSubject GameNotifier) *Team {
	return &Team{name: name, ip: ip, backlog: DefaultBacklog(), scoreSubject: scoreSubject}
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
	t.backlog.Done(userStoryIds...)
}

func (t *Team) CompleteIteration() {
	currentScore := t.backlog.Score()
	t.iterationScores = append(t.iterationScores, currentScore)
	if t.scoreSubject != nil {
		t.scoreSubject.NotifyScore(t.name, currentScore)
	}
}

func (t *Team) IterationScores() []Score {
	return t.iterationScores
}
