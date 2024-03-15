package calculator

import (
	. "elephant_carpaccio/domain/money"
)

type OrderGenerator struct {
	generatorRandom OrderRandom
	state           *State
	discount        *Discount
	withoutDecimals bool
}

func NewOrderGenerator(generatorRandom OrderRandom) *OrderGenerator {
	return &OrderGenerator{generatorRandom: generatorRandom}
}

func (og *OrderGenerator) WithState(state *State) *OrderGenerator {
	og.state = state
	return og
}

func (og *OrderGenerator) WithDiscount(discount *Discount) *OrderGenerator {
	og.discount = discount
	return og
}

func (og *OrderGenerator) WithoutDecimals(withoutDecimals bool) *OrderGenerator {
	og.withoutDecimals = withoutDecimals
	return og
}

func (og *OrderGenerator) GenerateOrder() Order {
	if og.state == nil {
		og.state = og.pickState()
	}
	if og.discount == nil {
		og.discount = og.pickDiscount()
	}
	nbItems := og.pickNbItems()
	itemPrice := og.pickItemPrice(og.discount, nbItems)
	return NewOrder(nbItems, itemPrice, og.state)
}

func (og *OrderGenerator) pickState() *State {
	return og.generatorRandom.RandState()
}

func (og *OrderGenerator) pickDiscount() *Discount {
	return og.generatorRandom.RandDiscountLevel()
}

func (og *OrderGenerator) pickNbItems() Decimal {
	if og.withoutDecimals {
		return og.generatorRandom.RandDecimalWithoutDecimals(Decimal(1), Decimal(10000))
	}
	return og.generatorRandom.RandDecimal(Decimal(1), Decimal(10000))
}

func (og *OrderGenerator) pickItemPrice(discount *Discount, nbItems Decimal) Dollar {
	minPrice, maxPrice := discount.AmountRange()
	minItemPrice := minPrice.Divide(nbItems)
	maxItemPrice := maxPrice.Divide(nbItems)
	itemPrice := og.generatorRandom.RandDollar(minItemPrice, maxItemPrice)
	return itemPrice
}
