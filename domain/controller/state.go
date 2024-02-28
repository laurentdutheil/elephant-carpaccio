package controller

import "elephant_carpaccio/domain/money"

type State struct {
	Label   string
	TaxRate money.Percent
}

func (s State) ApplyTax(amount money.Dollar) money.Dollar {
	return amount.Add(s.TaxRate.ApplyTo(amount))
}

func (s State) ComputeTax(amount money.Dollar) money.Dollar {
	return s.TaxRate.ApplyTo(amount)
}

type stateCode uint8

const (
	UT stateCode = iota
	NV
	TX
	AL
	CA

	numberOfStates
)

func (s stateCode) State() *State {
	return &AllStates()[s]
}

func AllStates() []State {
	return []State{
		{"UT", money.NewPercent(685)},
		{"NV", money.NewPercent(800)},
		{"TX", money.NewPercent(625)},
		{"AL", money.NewPercent(400)},
		{"CA", money.NewPercent(825)},
	}
}

func StateOf(value int) *State {
	if value < int(numberOfStates) {
		return stateCode(value).State()
	}
	return nil
}
