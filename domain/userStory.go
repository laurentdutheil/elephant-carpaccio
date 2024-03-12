package domain

import "elephant_carpaccio/domain/money"

type StoryId string

type UserStory struct {
	Id                      StoryId
	Description             string
	pointEstimation         int
	businessValueEstimation money.Dollar
	riskEstimation          int
	iterationEstimation     uint8
	done                    bool
	doneInIteration         uint8
}

func (u *UserStory) Done(currentIteration uint8) {
	if !u.done {
		u.done = true
		u.doneInIteration = currentIteration
	}
}

func (u *UserStory) IsDone() bool {
	return u.done
}

type UserStoryBuilder struct {
	userStory *UserStory
}

func NewUserStoryBuilder(id StoryId) *UserStoryBuilder {
	userStory := &UserStory{
		Id:                      id,
		pointEstimation:         1,
		businessValueEstimation: money.NewDollar(money.Decimal(0)),
	}
	b := &UserStoryBuilder{userStory: userStory}
	return b
}

func (b *UserStoryBuilder) Description(description string) *UserStoryBuilder {
	b.userStory.Description = description
	return b
}

func (b *UserStoryBuilder) BusinessValueEstimation(businessValueEstimation money.Dollar) *UserStoryBuilder {
	b.userStory.businessValueEstimation = businessValueEstimation
	return b
}

func (b *UserStoryBuilder) RiskEstimation(riskEstimation int) *UserStoryBuilder {
	b.userStory.riskEstimation = riskEstimation
	return b
}

func (b *UserStoryBuilder) IterationEstimation(iterationEstimation uint8) *UserStoryBuilder {
	b.userStory.iterationEstimation = iterationEstimation
	return b
}

func (b *UserStoryBuilder) Build() UserStory {
	return *b.userStory
}
