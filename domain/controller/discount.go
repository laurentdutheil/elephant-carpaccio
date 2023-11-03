package controller

import "sort"

type Discount struct {
	amount Dollar
	rate   float64
}

type Discounts []Discount

func (c Discounts) ComputeDiscountValue(orderValue Dollar) Dollar {
	orderedDiscounts := c.reversedSortDiscounts()
	for _, d := range orderedDiscounts {
		if orderValue.GreaterOrEqual(d.amount) {
			return orderValue.Multiply(d.rate)
		}
	}
	return NewDollar(0)
}

func (c Discounts) reversedSortDiscounts() Discounts {
	var orderedDiscounts = make(Discounts, len(c))
	println(copy(orderedDiscounts, c))
	sort.Slice(orderedDiscounts, func(i, j int) bool {
		return orderedDiscounts[i].amount.GreaterOrEqual(orderedDiscounts[j].amount)
	})
	return orderedDiscounts
}
