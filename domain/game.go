package domain

type Game struct {
	teams          []*Team
	scoreObservers map[string]GameObserver
}

func NewGame() *Game {
	observers := map[string]GameObserver{}
	return &Game{scoreObservers: observers}
}

func (g *Game) Register(teamName string, ip string) {
	if teamName != "" {
		g.teams = append(g.teams, NewTeam(teamName, ip, g))
		g.NotifyRegistration(teamName)
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

func (g *Game) AddGameObserver(observer GameObserver) {
	g.scoreObservers[observer.Id()] = observer
}

func (g *Game) RemoveGameObserver(id string) {
	delete(g.scoreObservers, id)
}

func (g *Game) NotifyScore(teamName string, newIterationScore Score) {
	for _, observer := range g.scoreObservers {
		observer.UpdateScore(teamName, newIterationScore)
	}
}

func (g *Game) NotifyRegistration(teamName string) {
	for _, observer := range g.scoreObservers {
		observer.AddRegistration(teamName)
	}
}

func (g *Game) NbGameObservers() int {
	return len(g.scoreObservers)
}
