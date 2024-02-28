package money

import "fmt"

type Percent struct {
	Decimal
}

func NewPercent(decimal Decimal) Percent {
	return Percent{Decimal: decimal}
}

func (p Percent) ApplyTo(amount Dollar) Dollar {
	return amount.Multiply(p.Decimal).Divide(10000)
}

func (p Percent) String() string {
	return fmt.Sprintf("%s%%", p.Decimal)
}
