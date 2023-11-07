package controller

func Compute(numberOfItems Decimal, itemPrice Dollar, stateCode StateCode) Receipt {
	state := stateCode.State()
	orderValue := itemPrice.Multiply(numberOfItems)
	discountValue := ComputeDiscount(orderValue)
	taxableValue := orderValue.Minus(discountValue)
	return Receipt{
		OrderValue:                orderValue,
		DiscountValue:             discountValue,
		TaxValue:                  state.ComputeTax(taxableValue),
		TotalPriceWithoutDiscount: state.ApplyTax(orderValue),
		TotalPrice:                state.ApplyTax(taxableValue),
	}
}

type Receipt struct {
	OrderValue                Dollar
	DiscountValue             Dollar
	TaxValue                  Dollar
	TotalPriceWithoutDiscount Dollar
	TotalPrice                Dollar
}
