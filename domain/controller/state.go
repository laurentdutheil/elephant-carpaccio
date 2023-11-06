package controller

type StateCode string

type State struct {
	stateCode StateCode
	taxRate   Percent
}

type States []State

func (s States) TaxRateOf(stateCode StateCode) Percent {
	for _, state := range s {
		if state.stateCode == stateCode {
			return state.taxRate
		}
	}
	return Percent(0)
}
