package domain

type Team struct {
	name    string
	backlog []UserStory
}

type UserStory struct {
	description string
	valuePoint  int
	done        bool
}

func NewTeam(name string) *Team {
	return &Team{name: name, backlog: defaultBacklog()}
}

func defaultBacklog() []UserStory {
	return []UserStory{
		{description: "Hello World", valuePoint: 1},
		{description: "Can fill parameters", valuePoint: 1},
		{description: "Compute order value without tax", valuePoint: 1},
		{description: "Can handle float for 'number of items' AND 'price by item'", valuePoint: 1},
		{description: "Tax for UT", valuePoint: 1},
		{description: "Tax for NV", valuePoint: 1},
		{description: "Tax for TX", valuePoint: 1},
		{description: "Tax for AL", valuePoint: 1},
		{description: "Tax for CA", valuePoint: 1},
		{description: "15% Discount", valuePoint: 1},
		{description: "10% Discount", valuePoint: 1},
		{description: "7% Discount", valuePoint: 1},
		{description: "5% Discount", valuePoint: 1},
		{description: "3% Discount", valuePoint: 1},
		{description: "Can handle rounding for result (two digits after the decimal point)", valuePoint: 1},
		{description: "Prompts are clear. Display currency", valuePoint: 1},
		{description: "Display details (order value, tax, discount", valuePoint: 1},
		{description: "Do not have to re-launch the application for each test", valuePoint: 1},
	}
}

func (t Team) Name() string {
	return t.name
}

func (t Team) Backlog() []UserStory {
	return t.backlog
}
