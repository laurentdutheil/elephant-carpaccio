package controller

type StateCode string

type State struct {
	stateCode StateCode
	taxRate   float64
}

type States []State

func (s States) TaxRateOf(stateCode StateCode) float64 {
	for _, state := range s {
		if state.stateCode == stateCode {
			return state.taxRate
		}
	}
	return 0
}
