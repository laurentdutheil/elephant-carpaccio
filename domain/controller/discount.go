package controller

type Discount struct {
	minAmount Dollar
	maxAmount Dollar
	Rate      Percent
}

func (d Discount) applyTo(amount Dollar) Dollar {
	return d.Rate.ApplyTo(amount)
}

func (d Discount) AmountRange() (minAmount Dollar, maxAmount Dollar) {
	return d.minAmount, d.maxAmount
}

type discountLevel uint8

const (
	No discountLevel = iota
	ThreePercent
	FivePercent
	SevenPercent
	TenPercent
	FifteenPercent

	numberOfDiscounts
)

func (l discountLevel) Discount() *Discount {
	return &AllDiscounts()[l]
}

func AllDiscounts() []Discount {
	return []Discount{
		{NewDollar(0), NewDollar(100000), NewPercent(0)},
		{NewDollar(100000), NewDollar(500000), NewPercent(300)},
		{NewDollar(500000), NewDollar(700000), NewPercent(500)},
		{NewDollar(700000), NewDollar(1000000), NewPercent(700)},
		{NewDollar(1000000), NewDollar(5000000), NewPercent(1000)},
		{NewDollar(5000000), NewDollar(100000000), NewPercent(1500)},
	}
}

func DiscountOf(value int) *Discount {
	if value < int(numberOfDiscounts) {
		return discountLevel(value).Discount()
	}
	return nil
}

func ComputeDiscount(amount Dollar) (Dollar, *Discount) {
	discount := findDiscount(amount)
	return discount.applyTo(amount), discount
}

func findDiscount(amount Dollar) *Discount {
	allDiscounts := AllDiscounts()
	for i := len(allDiscounts) - 1; i > 0; i-- {
		d := allDiscounts[i]
		if amount.GreaterOrEqual(d.minAmount) {
			return &d
		}
	}
	return No.Discount()
}
