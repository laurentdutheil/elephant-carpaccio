package controller

import (
	"fmt"
)

type Amount int64

type Dollar struct {
	amount Amount
}

func NewDollar(amountInCents Amount) Dollar {
	return Dollar{amount: amountInCents}
}

func (d Dollar) ApplyTaxe(taxRate float64) Dollar {
	return d.Add(d.Multiply(taxRate))
}

func (d Dollar) Multiply(mul float64) Dollar {
	r := mul * float64(d.amount)
	return NewDollar(Amount(r))
}

func (d Dollar) Divide(div float64) Dollar {
	r := float64(d.amount) / div
	return NewDollar(Amount(r))

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
