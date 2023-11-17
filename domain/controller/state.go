package controller

type State struct {
	Label   string
	TaxRate Percent
}

func (s State) ApplyTax(amount Dollar) Dollar {
	return amount.Add(s.TaxRate.ApplyTo(amount))
}

func (s State) ComputeTax(amount Dollar) Dollar {
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
		{"UT", NewPercent(685)},
		{"NV", NewPercent(800)},
		{"TX", NewPercent(625)},
		{"AL", NewPercent(400)},
		{"CA", NewPercent(825)},
	}
}

func StateOf(value int) *State {
	if value < int(numberOfStates) {
		return stateCode(value).State()
	}
	return nil
}
