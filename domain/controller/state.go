package controller

type State struct {
	Label   string
	TaxRate Percent
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
		{"UT", NewPercent(685)},
		{"NV", NewPercent(800)},
		{"TX", NewPercent(625)},
		{"AL", NewPercent(400)},
		{"CA", NewPercent(825)},
	}[s]
}

func (s State) ApplyTax(amount Dollar) Dollar {
	return amount.Add(s.TaxRate.ApplyTo(amount))
}

func (s State) ComputeTax(amount Dollar) Dollar {
	return s.TaxRate.ApplyTo(amount)
}
