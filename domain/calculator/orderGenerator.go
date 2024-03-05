package calculator

import (
	"elephant_carpaccio/domain/money"
)

type OrderGenerator struct {
	generatorRandom OrderRandom
}

func NewOrderGenerator(generatorRandom OrderRandom) *OrderGenerator {
	return &OrderGenerator{generatorRandom: generatorRandom}
}

func (og OrderGenerator) GenerateOrder(discount *Discount, state *State) Order {
	if state == nil {
		state = og.pickState()
	}
	if discount == nil {
		discount = og.pickDiscount()
	}
	nbItems := og.pickNbItems()
	itemPrice := og.pickItemPrice(discount, nbItems)
	return NewOrder(nbItems, itemPrice, state)
}

func (og OrderGenerator) pickState() *State {
	return og.generatorRandom.RandState()
}

func (og OrderGenerator) pickDiscount() *Discount {
	return og.generatorRandom.RandDiscountLevel()
}

func (og OrderGenerator) pickNbItems() money.Decimal {
	return og.generatorRandom.RandDecimal(money.Decimal(1), money.Decimal(10000))
}

func (og OrderGenerator) pickItemPrice(discount *Discount, nbItems money.Decimal) money.Dollar {
	minPrice, maxPrice := discount.AmountRange()
	minItemPrice := minPrice.Divide(nbItems)
	maxItemPrice := maxPrice.Divide(nbItems)
	itemPrice := og.generatorRandom.RandDollar(minItemPrice, maxItemPrice)
	return itemPrice
}
