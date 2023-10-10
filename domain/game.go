package domain

type Game struct {
	teams           []*Team
	iterationLogger *IterationLogger
}

func NewGame() *Game {
	return &Game{iterationLogger: NewIterationLogger()}
}

func (g *Game) Register(teamName string) {
	g.teams = append(g.teams, NewTeam(teamName))
}

func (g *Game) Teams() []*Team {
	return g.teams
}

func (g *Game) LogIteration() {
	for _, team := range g.teams {
		g.iterationLogger.LogIterationScore(team)
	}
}

func (g *Game) PrintBoard() string {
	return g.iterationLogger.String()
}
