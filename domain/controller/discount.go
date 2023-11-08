package controller

import "math"

type Discount struct {
	Amount Dollar
	Rate   Percent
}

type DiscountLevel int

const (
	NoDiscount DiscountLevel = iota
	ThreePercentDiscount
	FivePercentDiscount
	SevenPercentDiscount
	TenPercentDiscount
	FifteenPercentDiscount
	_numberOfDiscounts
)

func (l DiscountLevel) Discount() Discount {
	return []Discount{
		{NewDollar(0), NewPercent(0)},
		{NewDollar(100000), NewPercent(300)},
		{NewDollar(500000), NewPercent(500)},
		{NewDollar(700000), NewPercent(700)},
		{NewDollar(1000000), NewPercent(1000)},
		{NewDollar(5000000), NewPercent(1500)},
		{NewDollar(math.MaxInt64), NewPercent(0)},
	}[l]
}

func (l DiscountLevel) AmountRange() (minAmount Dollar, maxAmount Dollar) {
	minAmount = l.Discount().Amount
	maxAmount = (l + 1).Discount().Amount
	return
}

func ComputeDiscount(amount Dollar) Dollar {
	var discount Dollar
	for discountLevel := _numberOfDiscounts - 1; discountLevel >= NoDiscount; discountLevel-- {
		d := discountLevel.Discount()
		if amount.GreaterOrEqual(d.Amount) {
			discount = d.Rate.ApplyTo(amount)
			break
		}
	}
	return discount
}
