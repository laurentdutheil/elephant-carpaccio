package controller

import (
	"elephant_carpaccio/domain/money"
	"math/rand"
)

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

func (og OrderGenerator) generateItemPrice(discount *Discount, nbItems money.Decimal) money.Dollar {
	minPrice, maxPrice := discount.AmountRange()
	minItemPrice := minPrice.Divide(nbItems)
	maxItemPrice := maxPrice.Divide(nbItems)
	itemAmount := og.generatorRandom.randDecimal(minItemPrice.AmountInCents(), maxItemPrice.AmountInCents())
	return money.NewDollar(itemAmount)
}

func (og OrderGenerator) pickDiscount() *Discount {
	return DiscountOf(og.generatorRandom.randDiscountLevel())
}

func (og OrderGenerator) pickState() *State {
	return og.generatorRandom.randState()
}

type GeneratorRandom interface {
	randDecimal(minAmount money.Decimal, maxAmount money.Decimal) money.Decimal
	randDiscountLevel() DiscountLevel
	randState() *State
}

type randInt63nFunc func(max int64) int64

type GeneratorRandomizer struct {
	randFunc randInt63nFunc
}

func NewDecimalRandomizer(randFunc randInt63nFunc) *GeneratorRandomizer {
	return &GeneratorRandomizer{randFunc: randFunc}
}

func (dr GeneratorRandomizer) randDecimal(minAmount money.Decimal, maxAmount money.Decimal) money.Decimal {
	rangeAmount := maxAmount - minAmount
	orderValue := money.Decimal(dr.randFunc(int64(rangeAmount))) + minAmount
	return orderValue
}

func (dr GeneratorRandomizer) randInt(max int) int {
	return int(dr.randFunc(int64(max)))
}

func (dr GeneratorRandomizer) randDiscountLevel() DiscountLevel {
	randInt := dr.randInt(int(numberOfDiscounts))
	return DiscountLevel(randInt)
}

func (dr GeneratorRandomizer) randState() *State {
	randInt := dr.randInt(int(numberOfStates))
	return StateOf(randInt)
}
