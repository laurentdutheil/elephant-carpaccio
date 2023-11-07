package controller

import "math"

type Decimal int64

func (d Decimal) Multiply(other Decimal) Decimal {
	return Decimal(math.Round(float64(d) * float64(other) * math.Pow10(-2)))
}

func (d Decimal) Divide(other Decimal) Decimal {
	return Decimal(math.Round(float64(d) / float64(other) * math.Pow10(2)))
}

type Percent Decimal

func (p Percent) ApplyTo(amount Dollar) Dollar {
	return amount.Multiply(Decimal(p)).Divide(10000)
}
