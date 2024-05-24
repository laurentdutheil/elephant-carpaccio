package calculator

import "elephant_carpaccio/domain/money"

type State struct {
	Label   string
	TaxRate money.Percent
}

func StateOf(stateCode StateCode) *State {
	return stateCode.State()
}

func (s State) ApplyTax(amount money.Dollar) money.Dollar {
	return amount.Add(s.TaxRate.ApplyTo(amount))
}

func (s State) ComputeTax(amount money.Dollar) money.Dollar {
	return s.TaxRate.ApplyTo(amount)
}

type StateCode uint8

const (
	UT StateCode = iota
	NV
	TX
	AL
	CA

	numberOfStates
)

func (s StateCode) State() *State {
	if s < numberOfStates {
		return &allStates[s]
	}
	return nil
}

var allStates = []State{
	{"UT", money.NewPercent(685)},
	{"NV", money.NewPercent(800)},
	{"TX", money.NewPercent(625)},
	{"AL", money.NewPercent(400)},
	{"CA", money.NewPercent(825)},
}
