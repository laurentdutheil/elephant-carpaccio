package domain

type UserStory struct {
	Id          StoryId
	Description string
	ValuePoint  Score
	Done        bool
}

type StoryId string

type Backlog []UserStory

func (b Backlog) Done(id StoryId) {
	for i, story := range b {
		if id == story.Id {
			b[i].Done = true
			break
		}
	}
}

func (b Backlog) Score() Score {
	score := Score{0}
	for _, story := range b {
		if story.Done {
			score = score.Add(story.ValuePoint)
		}
	}
	return score
}

func defaultBacklog() Backlog {
	return Backlog{
		{Id: "EC-001", Description: "Hello World", ValuePoint: Score{1}},
		{Id: "EC-002", Description: "Can fill parameters", ValuePoint: Score{1}},
		{Id: "EC-003", Description: "Compute order value without tax", ValuePoint: Score{1}},
		{Id: "EC-004", Description: "Can handle float for 'number of items' AND 'price by item'", ValuePoint: Score{1}},
		{Id: "EC-005", Description: "Tax for UT", ValuePoint: Score{1}},
		{Id: "EC-006", Description: "Tax for NV", ValuePoint: Score{1}},
		{Id: "EC-007", Description: "Tax for TX", ValuePoint: Score{1}},
		{Id: "EC-008", Description: "Tax for AL", ValuePoint: Score{1}},
		{Id: "EC-009", Description: "Tax for CA", ValuePoint: Score{1}},
		{Id: "EC-010", Description: "15% Discount", ValuePoint: Score{1}},
		{Id: "EC-011", Description: "10% Discount", ValuePoint: Score{1}},
		{Id: "EC-012", Description: "7% Discount", ValuePoint: Score{1}},
		{Id: "EC-013", Description: "5% Discount", ValuePoint: Score{1}},
		{Id: "EC-014", Description: "3% Discount", ValuePoint: Score{1}},
		{Id: "EC-015", Description: "Can handle rounding for result (two digits after the decimal point)", ValuePoint: Score{1}},
		{Id: "EC-016", Description: "Prompts are clear. Display currency", ValuePoint: Score{1}},
		{Id: "EC-017", Description: "Display details (order value, tax, discount", ValuePoint: Score{1}},
		{Id: "EC-018", Description: "Do not have to re-launch the application for each test", ValuePoint: Score{1}},
	}
}
