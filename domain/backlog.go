package domain

type UserStory struct {
	id          StoryId
	description string
	valuePoint  int
	done        bool
}

type StoryId string

type Backlog []UserStory

func (b Backlog) Done(id StoryId) {
	for i, story := range b {
		if id == story.id {
			b[i].done = true
			break
		}
	}
}

func (b Backlog) Score() int {
	score := 0
	for _, story := range b {
		if story.done {
			score += story.valuePoint
		}
	}
	return score

}

func defaultBacklog() Backlog {
	return Backlog{
		{id: "EC-001", description: "Hello World", valuePoint: 1},
		{id: "EC-002", description: "Can fill parameters", valuePoint: 1},
		{id: "EC-003", description: "Compute order value without tax", valuePoint: 1},
		{id: "EC-004", description: "Can handle float for 'number of items' AND 'price by item'", valuePoint: 1},
		{id: "EC-005", description: "Tax for UT", valuePoint: 1},
		{id: "EC-006", description: "Tax for NV", valuePoint: 1},
		{id: "EC-007", description: "Tax for TX", valuePoint: 1},
		{id: "EC-008", description: "Tax for AL", valuePoint: 1},
		{id: "EC-009", description: "Tax for CA", valuePoint: 1},
		{id: "EC-010", description: "15% Discount", valuePoint: 1},
		{id: "EC-011", description: "10% Discount", valuePoint: 1},
		{id: "EC-012", description: "7% Discount", valuePoint: 1},
		{id: "EC-013", description: "5% Discount", valuePoint: 1},
		{id: "EC-014", description: "3% Discount", valuePoint: 1},
		{id: "EC-015", description: "Can handle rounding for result (two digits after the decimal point)", valuePoint: 1},
		{id: "EC-016", description: "Prompts are clear. Display currency", valuePoint: 1},
		{id: "EC-017", description: "Display details (order value, tax, discount", valuePoint: 1},
		{id: "EC-018", description: "Do not have to re-launch the application for each test", valuePoint: 1},
	}
}
