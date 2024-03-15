package calculator

import (
	"math/rand"

	. "elephant_carpaccio/domain/money"
)

type OrderRandom interface {
	RandDecimal(min Decimal, max Decimal) Decimal
	RandDecimalWithoutDecimals(min Decimal, max Decimal) Decimal
	RandDollar(minAmount Dollar, maxAmount Dollar) Dollar
	RandDiscountLevel() *Discount
	RandState() *State
}

var RandIntn = rand.Intn
var RandInt63n = rand.Int63n

type OrderRandomizer struct {
}

func NewOrderRandomizer() *OrderRandomizer {
	return &OrderRandomizer{}
}

func (dr OrderRandomizer) RandDecimal(min Decimal, max Decimal) Decimal {
	rangeDecimal := max - min
	randomRange := Decimal(RandInt63n(int64(rangeDecimal)))
	return min + randomRange
}

func (dr OrderRandomizer) RandDecimalWithoutDecimals(min Decimal, max Decimal) Decimal {
	rangeDecimal := max - min
	randomRange := Decimal(RandInt63n(int64(rangeDecimal)))
	randomRange = randomRange.Divide(Decimal(10000)).Multiply(Decimal(10000))
	minWithoutDecimals := min.Divide(Decimal(10000)).Multiply(Decimal(10000))
	return minWithoutDecimals + randomRange
}

func (dr OrderRandomizer) RandDollar(minAmount Dollar, maxAmount Dollar) Dollar {
	rangeDollar := maxAmount.Minus(minAmount)
	randomRange := Decimal(RandInt63n(int64(rangeDollar.AmountInCents())))
	return minAmount.Add(NewDollar(randomRange))
}

func (dr OrderRandomizer) RandDiscountLevel() *Discount {
	randDiscountLevel := RandIntn(int(numberOfDiscounts))
	return DiscountOf(randDiscountLevel)
}

func (dr OrderRandomizer) RandState() *State {
	randStateCode := RandIntn(int(numberOfStates))
	return StateOf(randStateCode)
}
