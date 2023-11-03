package controller

var AllStates = States{
	{"UT", 0.0685},
	{"NV", 0.08},
	{"TX", 0.0625},
	{"AL", 0.04},
	{"CA", 0.0825},
}

var AllDiscounts = Discounts{
	{NewDollar(100000), 0.03},
	{NewDollar(500000), 0.05},
	{NewDollar(700000), 0.07},
	{NewDollar(1000000), 0.1},
	{NewDollar(5000000), 0.15},
}

func Compute(numberOfItems float64, itemPrice Dollar, stateCode StateCode) Receipt {
	taxRate := AllStates.TaxRateOf(stateCode)
	orderValue := itemPrice.Multiply(numberOfItems)
	discountValue := AllDiscounts.ComputeDiscountValue(orderValue)
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
