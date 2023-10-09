package domain

type Team struct {
	name string
}

func NewTeam(name string) *Team {
	return &Team{name: name}
}

func (t Team) Name() string {
	return t.name
}
