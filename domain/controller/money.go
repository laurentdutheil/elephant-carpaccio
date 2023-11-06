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

type Percent Decimal

func (p Percent) ApplyTo(other Dollar) Dollar {
	return other.Multiply(Decimal(p)).Divide(10000)
}

type Dollar struct {
	amount Decimal
}

func NewDollar(amountInCents Decimal) Dollar {
	return Dollar{amount: amountInCents}
}

func (d Dollar) ApplyTaxe(taxRate Percent) Dollar {
	return d.Add(taxRate.ApplyTo(d))
}

func (d Dollar) Multiply(mul Decimal) Dollar {
	r := d.amount.Multiply(mul)
	return NewDollar(r)
}

func (d Dollar) Divide(div Decimal) Dollar {
	r := d.amount.Divide(div)
	return NewDollar(r)
}

func (d Dollar) GreaterOrEqual(other Dollar) bool {
	return d.amount >= other.amount
}

func (d Dollar) Lower(other Dollar) bool {
	return d.amount < other.amount
}

func (d Dollar) Add(other Dollar) Dollar {
	return NewDollar(d.amount + other.amount)
}

func (d Dollar) Minus(other Dollar) Dollar {
	return NewDollar(d.amount - other.amount)
}

func (d Dollar) String() string {
	dollars := d.amount / 100
	cents := d.amount % 100
	return fmt.Sprintf("$%d.%02d", dollars, cents)
}
