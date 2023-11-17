package controller

import "math/rand"

type OrderGenerator struct {
	generatorRandom GeneratorRandom
}

func NewOrderGenerator(generatorRandom GeneratorRandom) *OrderGenerator {
	if generatorRandom == nil {
		generatorRandom = NewDecimalRandomizer(rand.Int63n)
	}
	return &OrderGenerator{generatorRandom: generatorRandom}
}

func (og OrderGenerator) GenerateOrder(discount *Discount, state *State) Order {
	if state == nil {
		state = og.pickState()
	}
	if discount == nil {
		discount = og.pickDiscount()
	}
	nbItems := og.generatorRandom.randDecimal(1, 10000)
	itemPrice := og.generateItemPrice(discount, nbItems)
	return NewOrder(nbItems, itemPrice, state)
}

func (og OrderGenerator) generateItemPrice(discount *Discount, nbItems Decimal) Dollar {
	minPrice, maxPrice := discount.AmountRange()
	minItemPrice := minPrice.Divide(nbItems)
	maxItemPrice := maxPrice.Divide(nbItems)
	itemAmount := og.generatorRandom.randDecimal(minItemPrice.amount, maxItemPrice.amount)
	return NewDollar(itemAmount)
}

func (og OrderGenerator) pickDiscount() *Discount {
	return DiscountOf(og.generatorRandom.randInt(int(numberOfDiscounts)))
}

func (og OrderGenerator) pickState() *State {
	return StateOf(og.generatorRandom.randInt(int(numberOfStates)))
}

type GeneratorRandom interface {
	randDecimal(minAmount Decimal, maxAmount Decimal) Decimal
	randInt(max int) int
}

type randInt63nFunc func(max int64) int64

type GeneratorRandomizer struct {
	randFunc randInt63nFunc
}

func NewDecimalRandomizer(randFunc randInt63nFunc) *GeneratorRandomizer {
	return &GeneratorRandomizer{randFunc: randFunc}
}

func (dr GeneratorRandomizer) randDecimal(minAmount Decimal, maxAmount Decimal) Decimal {
	rangeAmount := maxAmount - minAmount
	orderValue := Decimal(dr.randFunc(int64(rangeAmount))) + minAmount
	return orderValue
}

func (dr GeneratorRandomizer) randInt(max int) int {
	return int(dr.randFunc(int64(max)))
}
