package controller

import (
	"math/rand"
)

var AllStates = States{
	{"UT", Percent(685)},
	{"NV", Percent(800)},
	{"TX", Percent(625)},
	{"AL", Percent(400)},
	{"CA", Percent(825)},
}

func Compute(numberOfItems Decimal, itemPrice Dollar, stateCode StateCode) Receipt {
	taxRate := AllStates.TaxRateOf(stateCode)
	orderValue := itemPrice.Multiply(numberOfItems)
	discountValue := ComputeDiscountValue(orderValue)
	taxableValue := orderValue.Minus(discountValue)
	return Receipt{
		OrderValue:                orderValue,
		DiscountValue:             discountValue,
		TaxValue:                  taxRate.ApplyTo(taxableValue),
		TotalPriceWithoutDiscount: orderValue.ApplyTaxe(taxRate),
		TotalPrice:                taxableValue.ApplyTaxe(taxRate),
	}
}

type Receipt struct {
	OrderValue                Dollar
	DiscountValue             Dollar
	TaxValue                  Dollar
	TotalPriceWithoutDiscount Dollar
	TotalPrice                Dollar
}

func GenerateOrder(discountLevel DiscountLevel) (Decimal, Dollar) {
	nbItems := Decimal(rand.Int63n(1000)) + 1
	itemPrice := generateItemPrice(discountLevel, nbItems)
	return nbItems, itemPrice
}

func generateItemPrice(discountLevel DiscountLevel, nbItems Decimal) Dollar {
	minAmount, maxAmount := discountLevel.AmountRange()
	orderValue := randAmount(minAmount, maxAmount)
	itemPrice := NewDollar(orderValue).Divide(nbItems)
	return itemPrice
}

func randAmount(minAmount Dollar, maxAmount Dollar) Decimal {
	rangeAmount := maxAmount.Minus(minAmount)
	orderValue := Decimal(rand.Int63n(int64(rangeAmount.amount))) + minAmount.amount
	return orderValue
}
