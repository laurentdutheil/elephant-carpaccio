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
		{NewDollar(0), 0},
		{NewDollar(100000), 300},
		{NewDollar(500000), 500},
		{NewDollar(700000), 700},
		{NewDollar(1000000), 1000},
		{NewDollar(5000000), 1500},
		{NewDollar(math.MaxInt64), 0.0},
	}[l]
}

func (l DiscountLevel) AmountRange() (minAmount Dollar, maxAmount Dollar) {
	minAmount = l.Discount().Amount
	maxAmount = (l + 1).Discount().Amount
	return
}

func ComputeDiscountValue(orderValue Dollar) Dollar {
	var discountValue Dollar
	for discountLevel := _numberOfDiscounts - 1; discountLevel >= NoDiscount; discountLevel-- {
		d := discountLevel.Discount()
		if orderValue.GreaterOrEqual(d.Amount) {
			discountValue = d.Rate.ApplyTo(orderValue)
			break
		}
	}
	return discountValue
}
