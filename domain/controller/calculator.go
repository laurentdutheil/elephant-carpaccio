package controller

import (
	"math"
	"math/rand"
)

var AllStates = States{
	{"UT", 0.0685},
	{"NV", 0.08},
	{"TX", 0.0625},
	{"AL", 0.04},
	{"CA", 0.0825},
}

func Compute(numberOfItems float64, itemPrice Dollar, stateCode StateCode) Receipt {
	taxRate := AllStates.TaxRateOf(stateCode)
	orderValue := itemPrice.Multiply(numberOfItems)
	discountValue := ComputeDiscountValue(orderValue)
	taxableValue := orderValue.Minus(discountValue)
	return Receipt{
		OrderValue:                orderValue,
		DiscountValue:             discountValue,
		TaxValue:                  taxableValue.Multiply(taxRate),
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

func GenerateOrder(discountLevel DiscountLevel) (float64, Dollar) {
	nbItems := math.Trunc(rand.Float64()*math.Pow10(4)) * math.Pow10(-2)
	itemPrice := generateItemPrice(discountLevel, nbItems)
	return nbItems, itemPrice
}

func generateItemPrice(discountLevel DiscountLevel, nbItems float64) Dollar {
	minAmount, maxAmount := discountLevel.AmountRange()
	orderValue := randAmount(minAmount, maxAmount)
	itemPrice := NewDollar(orderValue).Divide(nbItems)
	return itemPrice
}

func randAmount(minAmount Dollar, maxAmount Dollar) Amount {
	rangeAmount := maxAmount.Minus(minAmount)
	orderValue := Amount(rand.Int63n(int64(rangeAmount.amount))) + minAmount.amount
	return orderValue
}
