package controller

import "math"

type Discount struct {
	Amount Dollar
	Rate   float64
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
		{NewDollar(0), 0.0},
		{NewDollar(100000), 0.03},
		{NewDollar(500000), 0.05},
		{NewDollar(700000), 0.07},
		{NewDollar(1000000), 0.1},
		{NewDollar(5000000), 0.15},
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
			discountValue = orderValue.Multiply(d.Rate)
			break
		}
	}
	return discountValue
}
