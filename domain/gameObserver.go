package domain

type GameNotifier interface {
	AddGameObserver(observer GameObserver)
	RemoveGameObserver(id string)
	NotifyScore(teamName string, newScore Score)
	NotifyRegistration(teamName string)
}

type GameObserver interface {
	Id() string
	UpdateScore(teamName string, newScore Score)
	AddRegistration(teamName string)
}

type GameSubject struct {
	scoreObservers map[string]GameObserver
}

func NewGameSubject() GameSubject {
	observers := map[string]GameObserver{}
	return GameSubject{scoreObservers: observers}
}

func (g *GameSubject) AddGameObserver(observer GameObserver) {
	g.scoreObservers[observer.Id()] = observer
}

func (g *GameSubject) RemoveGameObserver(id string) {
	delete(g.scoreObservers, id)
}

func (g *GameSubject) NotifyScore(teamName string, newIterationScore Score) {
	for _, observer := range g.scoreObservers {
		observer.UpdateScore(teamName, newIterationScore)
	}
}

func (g *GameSubject) NotifyRegistration(teamName string) {
	for _, observer := range g.scoreObservers {
		observer.AddRegistration(teamName)
	}
}

func (g *GameSubject) NbGameObservers() int {
	return len(g.scoreObservers)
}
