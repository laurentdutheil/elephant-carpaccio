package domain

import "fmt"

type Game struct {
	teams []*Team
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Register(teamName string) {
	g.teams = append(g.teams, NewTeam(teamName))
}

func (g *Game) Teams() []*Team {
	return g.teams
}

func (g *Game) LogIteration() {
	for _, team := range g.teams {
		team.LogIterationScore()
	}
}

func (g *Game) PrintBoard() string {
	result := ""
	for _, team := range g.teams {
		result += team.name + ": "
		result += fmt.Sprintln(team.IterationScores())
	}
	return result
}
