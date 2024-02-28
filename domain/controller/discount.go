package controller

import "elephant_carpaccio/domain/money"

type Discount struct {
	minAmount money.Dollar
	maxAmount money.Dollar
	Rate      money.Percent
}

func (d Discount) applyTo(amount money.Dollar) money.Dollar {
	return d.Rate.ApplyTo(amount)
}

func (d Discount) AmountRange() (minAmount money.Dollar, maxAmount money.Dollar) {
	return d.minAmount, d.maxAmount
}

type DiscountLevel uint8

const (
	No DiscountLevel = iota
	ThreePercent
	FivePercent
	SevenPercent
	TenPercent
	FifteenPercent

	numberOfDiscounts
)

func (l DiscountLevel) Discount() *Discount {
	return &AllDiscounts()[l]
}

func AllDiscounts() []Discount {
	return []Discount{
		{money.NewDollar(0), money.NewDollar(100000), money.NewPercent(0)},
		{money.NewDollar(100000), money.NewDollar(500000), money.NewPercent(300)},
		{money.NewDollar(500000), money.NewDollar(700000), money.NewPercent(500)},
		{money.NewDollar(700000), money.NewDollar(1000000), money.NewPercent(700)},
		{money.NewDollar(1000000), money.NewDollar(5000000), money.NewPercent(1000)},
		{money.NewDollar(5000000), money.NewDollar(100000000), money.NewPercent(1500)},
	}
}

func DiscountOf(level DiscountLevel) *Discount {
	if level < numberOfDiscounts {
		return level.Discount()
	}
	return nil
}

func ComputeDiscount(amount money.Dollar) (money.Dollar, *Discount) {
	discount := findDiscount(amount)
	return discount.applyTo(amount), discount
}

func findDiscount(amount money.Dollar) *Discount {
	allDiscounts := AllDiscounts()
	for i := len(allDiscounts) - 1; i > 0; i-- {
		d := allDiscounts[i]
		if amount.GreaterOrEqual(d.minAmount) {
			return &d
		}
	}
	return No.Discount()
}
