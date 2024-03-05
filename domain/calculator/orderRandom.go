package calculator

import (
	"elephant_carpaccio/domain/money"
	"math/rand"
)

type OrderRandom interface {
	RandDecimal(min money.Decimal, max money.Decimal) money.Decimal
	RandDollar(minAmount money.Dollar, maxAmount money.Dollar) money.Dollar
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

func (dr OrderRandomizer) RandDecimal(min money.Decimal, max money.Decimal) money.Decimal {
	rangeDecimal := max - min
	randomRange := money.Decimal(RandInt63n(int64(rangeDecimal)))
	return min + randomRange
}

func (dr OrderRandomizer) RandDollar(minAmount money.Dollar, maxAmount money.Dollar) money.Dollar {
	rangeDollar := maxAmount.Minus(minAmount)
	randomRange := money.Decimal(RandInt63n(int64(rangeDollar.AmountInCents())))
	return minAmount.Add(money.NewDollar(randomRange))
}

func (dr OrderRandomizer) RandDiscountLevel() *Discount {
	randDiscountLevel := RandIntn(int(numberOfDiscounts))
	return DiscountOf(randDiscountLevel)
}

func (dr OrderRandomizer) RandState() *State {
	randStateCode := RandIntn(int(numberOfStates))
	return StateOf(randStateCode)
}
