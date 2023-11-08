package controller

import (
	"fmt"
	"math"
)

type Decimal int64

func (d Decimal) Multiply(other Decimal) Decimal {
	return Decimal(math.Round(float64(d) * float64(other) * math.Pow10(-2)))
}

func (d Decimal) Divide(other Decimal) Decimal {
	return Decimal(math.Round(float64(d) / float64(other) * math.Pow10(2)))
}

func (d Decimal) String() string {
	units := d / 100
	decimals := d % 100
	return fmt.Sprintf("%d.%02d", units, decimals)
}

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
