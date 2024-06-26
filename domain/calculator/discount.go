package calculator

import "elephant_carpaccio/domain/money"

type Discount struct {
	minAmount money.Dollar
	Rate      money.Percent
}

func DiscountOf(level DiscountLevel) *Discount {
	return level.Discount()
}

func FindDiscount(amount money.Dollar) *Discount {
	for i := len(allDiscounts) - 1; i > 0; i-- {
		d := allDiscounts[i]
		if amount.GreaterOrEqual(d.minAmount) {
			return &d
		}
	}
	return No.Discount()
}

func (d Discount) AmountRange() (minAmount money.Dollar, maxAmount money.Dollar) {
	return d.minAmount, d.findMaxAmount()
}

func (d Discount) ComputeDiscount(amount money.Dollar) money.Dollar {
	return d.Rate.ApplyTo(amount)
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
	if l < numberOfDiscounts {
		return &allDiscounts[l]
	}
	return nil
}

var allDiscounts = []Discount{
	{money.NewDollar(0), money.NewPercent(0)},
	{money.NewDollar(100000), money.NewPercent(300)},
	{money.NewDollar(500000), money.NewPercent(500)},
	{money.NewDollar(700000), money.NewPercent(700)},
	{money.NewDollar(1000000), money.NewPercent(1000)},
	{money.NewDollar(5000000), money.NewPercent(1500)},
}

func (d Discount) findMaxAmount() money.Dollar {
	nextDiscount := d.findNextDiscount()
	if nextDiscount == nil {
		return money.NewDollar(100000000)
	}
	return nextDiscount.minAmount
}

func (d Discount) findNextDiscount() *Discount {
	dl := No
	for ; *(dl.Discount()) != d; dl++ {
	}
	return (dl + 1).Discount()
}
