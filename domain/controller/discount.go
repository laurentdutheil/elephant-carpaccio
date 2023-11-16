package controller

type Discount struct {
	Amount Dollar
	Rate   Percent
}

func (d Discount) applyTo(amount Dollar) Dollar {
	return d.Rate.ApplyTo(amount)
}

type DiscountLevel int

const (
	NoDiscount DiscountLevel = iota
	ThreePercentDiscount
	FivePercentDiscount
	SevenPercentDiscount
	TenPercentDiscount
	FifteenPercentDiscount

	NumberOfDiscounts
)

func (l DiscountLevel) Discount() Discount {
	return []Discount{
		{NewDollar(0), NewPercent(0)},
		{NewDollar(100000), NewPercent(300)},
		{NewDollar(500000), NewPercent(500)},
		{NewDollar(700000), NewPercent(700)},
		{NewDollar(1000000), NewPercent(1000)},
		{NewDollar(5000000), NewPercent(1500)},
		{NewDollar(100000000), NewPercent(0)},
	}[l]
}

func (l DiscountLevel) AmountRange() (minAmount Dollar, maxAmount Dollar) {
	minAmount = l.Discount().Amount
	maxAmount = (l + 1).Discount().Amount
	return
}

func ComputeDiscount(amount Dollar) (Dollar, Discount) {
	discount := findDiscount(amount)
	return discount.applyTo(amount), discount
}

func findDiscount(amount Dollar) Discount {
	for discountLevel := NumberOfDiscounts - 1; discountLevel > NoDiscount; discountLevel-- {
		d := discountLevel.Discount()
		if amount.GreaterOrEqual(d.Amount) {
			return d
		}
	}
	return NoDiscount.Discount()
}
