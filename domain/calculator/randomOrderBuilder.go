package calculator

import (
	. "elephant_carpaccio/domain/money"
)

type RandomOrderBuilder struct {
	orderRandom     OrderRandom
	state           *State
	discount        *Discount
	withoutDecimals bool
}

func NewRandomOrderBuilder(orderRandom OrderRandom) *RandomOrderBuilder {
	return &RandomOrderBuilder{orderRandom: orderRandom}
}

func (og *RandomOrderBuilder) WithState(state *State) *RandomOrderBuilder {
	og.state = state
	return og
}

func (og *RandomOrderBuilder) WithDiscount(discount *Discount) *RandomOrderBuilder {
	og.discount = discount
	return og
}

func (og *RandomOrderBuilder) WithoutDecimals(withoutDecimals bool) *RandomOrderBuilder {
	og.withoutDecimals = withoutDecimals
	return og
}

func (og *RandomOrderBuilder) Build() Order {
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

func (og *RandomOrderBuilder) pickState() *State {
	return og.orderRandom.RandState()
}

func (og *RandomOrderBuilder) pickDiscount() *Discount {
	return og.orderRandom.RandDiscountLevel()
}

func (og *RandomOrderBuilder) pickNbItems() Decimal {
	if og.withoutDecimals {
		return og.orderRandom.RandDecimalWithoutDecimals(Decimal(1), Decimal(10000))
	}
	return og.orderRandom.RandDecimal(Decimal(1), Decimal(10000))
}

func (og *RandomOrderBuilder) pickItemPrice(discount *Discount, nbItems Decimal) Dollar {
	minPrice, maxPrice := discount.AmountRange()
	minItemPrice := minPrice.Divide(nbItems)
	maxItemPrice := maxPrice.Divide(nbItems)
	itemPrice := og.orderRandom.RandDollar(minItemPrice, maxItemPrice)
	return itemPrice
}
