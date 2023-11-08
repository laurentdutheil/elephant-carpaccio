package controller

import (
	"fmt"
)

type Dollar struct {
	amount Decimal
}

func NewDollar(amountInCents Decimal) Dollar {
	return Dollar{amount: amountInCents}
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
	return fmt.Sprintf("$%v", d.amount)
}
