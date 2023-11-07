package controller

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
