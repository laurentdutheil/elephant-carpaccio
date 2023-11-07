package controller

type State struct {
	stateCode string
	taxRate   Percent
}

type StateCode int

const (
	UT StateCode = iota
	NV
	TX
	AL
	CA
)

func (s StateCode) State() State {
	return []State{
		{"UT", Percent(685)},
		{"NV", Percent(800)},
		{"TX", Percent(625)},
		{"AL", Percent(400)},
		{"CA", Percent(825)},
	}[s]
}

func (s State) ApplyTax(amount Dollar) Dollar {
	return amount.Add(s.taxRate.ApplyTo(amount))
}

func (s State) ComputeTax(amount Dollar) Dollar {
	return s.taxRate.ApplyTo(amount)
}
