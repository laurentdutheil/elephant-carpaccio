package domain

type UserStory struct {
	Id          StoryId
	Description string
	Score       Score
	Done        bool
}

type StoryId string

type Backlog []UserStory

func (b Backlog) Done(userStoryIds ...StoryId) {
	for i, story := range b {
		if contains(userStoryIds, story.Id) {
			b[i].Done = true
		}
	}
}

func contains(s []StoryId, e StoryId) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (b Backlog) Score() Score {
	score := Score{0}
	for _, story := range b {
		if story.Done {
			score = score.Add(story.Score)
		}
	}
	return score
}

func DefaultBacklog() Backlog {
	return Backlog{
		{Id: "EC-001", Description: "Hello World", Score: Score{1}},
		{Id: "EC-002", Description: "Can fill parameters", Score: Score{1}},
		{Id: "EC-003", Description: "Compute order value without tax", Score: Score{1}},
		{Id: "EC-004", Description: "Can handle float for 'number of items' AND 'price by item'", Score: Score{1}},
		{Id: "EC-005", Description: "Tax for UT", Score: Score{1}},
		{Id: "EC-006", Description: "Tax for NV", Score: Score{1}},
		{Id: "EC-007", Description: "Tax for TX", Score: Score{1}},
		{Id: "EC-008", Description: "Tax for AL", Score: Score{1}},
		{Id: "EC-009", Description: "Tax for CA", Score: Score{1}},
		{Id: "EC-010", Description: "15% Discount", Score: Score{1}},
		{Id: "EC-011", Description: "10% Discount", Score: Score{1}},
		{Id: "EC-012", Description: "7% Discount", Score: Score{1}},
		{Id: "EC-013", Description: "5% Discount", Score: Score{1}},
		{Id: "EC-014", Description: "3% Discount", Score: Score{1}},
		{Id: "EC-015", Description: "Can handle rounding for result (two digits after the decimal point)", Score: Score{1}},
		{Id: "EC-016", Description: "Prompts are clear. Display currency", Score: Score{1}},
		{Id: "EC-017", Description: "Display details (order value, tax, discount", Score: Score{1}},
		{Id: "EC-018", Description: "Do not have to re-launch the application for each test", Score: Score{1}},
	}
}
