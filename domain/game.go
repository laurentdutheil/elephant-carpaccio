package domain

type Game struct {
	teams          []*Team
	scoreObservers map[string]ScoreObserver
}

func NewGame() *Game {
	observers := map[string]ScoreObserver{}
	return &Game{scoreObservers: observers}
}

func (g *Game) Register(teamName string) {
	if teamName != "" {
		g.teams = append(g.teams, NewTeam(teamName, g))
	}
}

func (g *Game) Teams() []*Team {
	return g.teams
}

func (g *Game) FindTeamByName(teamName string) *Team {
	var selectedTeam *Team
	for _, team := range g.Teams() {
		if team.Name() == teamName {
			selectedTeam = team
			break
		}
	}
	return selectedTeam
}

func (g *Game) AddScoreObserver(observer ScoreObserver) {
	g.scoreObservers[observer.Id()] = observer
}

func (g *Game) RemoveScoreObserver(id string) {
	delete(g.scoreObservers, id)
}

func (g *Game) NotifyAll(teamName string, newIterationScore Score) {
	for _, observer := range g.scoreObservers {
		observer.Update(teamName, newIterationScore)
	}
}
