package domain

type GameSubject interface {
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
