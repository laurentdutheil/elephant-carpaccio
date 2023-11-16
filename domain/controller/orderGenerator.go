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

func (og OrderGenerator) GenerateOrder(discountLevel DiscountLevel, stateCode StateCode) Order {
	nbItems := og.generatorRandom.randDecimal(1, 10000)
	itemPrice := og.generateItemPrice(discountLevel, nbItems)
	return NewOrder(nbItems, itemPrice, stateCode)
}

func (og OrderGenerator) generateItemPrice(discountLevel DiscountLevel, nbItems Decimal) Dollar {
	minPrice, maxPrice := discountLevel.AmountRange()
	minItemPrice := minPrice.Divide(nbItems)
	maxItemPrice := maxPrice.Divide(nbItems)
	itemAmount := og.generatorRandom.randDecimal(minItemPrice.amount, maxItemPrice.amount)
	return NewDollar(itemAmount)
}

func (og OrderGenerator) PickDiscountLevel() DiscountLevel {
	return DiscountLevel(og.generatorRandom.randInt(int(NumberOfDiscounts)))
}

func (og OrderGenerator) PickStateCode() StateCode {
	return StateCode(og.generatorRandom.randInt(int(NumberOfStates)))
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
