package domain

type Game struct {
	teams []*Team
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Register(teamName string) {
	if teamName != "" {
		g.teams = append(g.teams, NewTeam(teamName, nil))
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
