package domain

import "elephant_carpaccio/domain/money"

type StoryId string

func (id StoryId) IsIn(ids []StoryId) bool {
	for _, i := range ids {
		if i == id {
			return true
		}
	}
	return false
}

type UserStory struct {
	Id                      StoryId
	Description             string
	PointEstimation         int
	BusinessValueEstimation money.Dollar
	RiskEstimation          int
	IterationEstimation     uint8
	done                    bool
	doneInIteration         uint8
}

func (u *UserStory) Done(inIteration uint8) {
	if !u.done {
		u.done = true
		u.doneInIteration = inIteration
	}
}

func (u *UserStory) IsDone() bool {
	return u.done
}

type UserStoryBuilder struct {
	userStory *UserStory
}

func NewUserStoryBuilder(id StoryId) *UserStoryBuilder {
	return &UserStoryBuilder{
		userStory: &UserStory{
			Id:              id,
			PointEstimation: 1,
		},
	}
}

func (b *UserStoryBuilder) Description(description string) *UserStoryBuilder {
	b.userStory.Description = description
	return b
}

func (b *UserStoryBuilder) BusinessValueEstimation(businessValueEstimation money.Dollar) *UserStoryBuilder {
	b.userStory.BusinessValueEstimation = businessValueEstimation
	return b
}

func (b *UserStoryBuilder) RiskEstimation(riskEstimation int) *UserStoryBuilder {
	b.userStory.RiskEstimation = riskEstimation
	return b
}

func (b *UserStoryBuilder) IterationEstimation(iterationEstimation uint8) *UserStoryBuilder {
	b.userStory.IterationEstimation = iterationEstimation
	return b
}

func (b *UserStoryBuilder) Build() UserStory {
	return *b.userStory
}
