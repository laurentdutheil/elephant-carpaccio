package controller

type OrderGenerator struct {
	randomizer DecimalRandom
}

func NewOrderGenerator(randomizer DecimalRandom) *OrderGenerator {
	return &OrderGenerator{randomizer: randomizer}
}

func (og OrderGenerator) GenerateOrder(discountLevel DiscountLevel, stateCode StateCode) Order {
	nbItems := og.randomizer.randDecimal(1, 10000)
	itemPrice := og.generateItemPrice(discountLevel, nbItems)
	return NewOrder(nbItems, itemPrice, stateCode)
}

func (og OrderGenerator) generateItemPrice(discountLevel DiscountLevel, nbItems Decimal) Dollar {
	minPrice, maxPrice := discountLevel.AmountRange()
	minItemPrice := minPrice.Divide(nbItems)
	maxItemPrice := maxPrice.Divide(nbItems)
	itemAmount := og.randomizer.randDecimal(minItemPrice.amount, maxItemPrice.amount)
	return NewDollar(itemAmount)
}

type DecimalRandom interface {
	randDecimal(minAmount Decimal, maxAmount Decimal) Decimal
}

type randInt63nFunc func(max int64) int64

type DecimalRandomizer struct {
	randFunc randInt63nFunc
}

func NewDecimalRandomizer(randFunc randInt63nFunc) *DecimalRandomizer {
	return &DecimalRandomizer{randFunc: randFunc}
}

func (dr DecimalRandomizer) randDecimal(minAmount Decimal, maxAmount Decimal) Decimal {
	rangeAmount := maxAmount - minAmount
	orderValue := Decimal(dr.randFunc(int64(rangeAmount))) + minAmount
	return orderValue
}
