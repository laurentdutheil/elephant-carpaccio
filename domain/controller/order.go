package controller

type Order struct {
	NumberOfItems Decimal
	ItemPrice     Dollar
	State         State
}

func NewOrder(numberOfItems Decimal, itemPrice Dollar, stateCode StateCode) Order {
	return Order{NumberOfItems: numberOfItems, ItemPrice: itemPrice, State: stateCode.State()}
}

func (o Order) Compute() Receipt {
	orderValue := o.ItemPrice.Multiply(o.NumberOfItems)
	discountValue := ComputeDiscount(orderValue)
	taxableValue := orderValue.Minus(discountValue)
	return Receipt{
		OrderValue:                orderValue,
		DiscountValue:             discountValue,
		TaxValue:                  o.State.ComputeTax(taxableValue),
		TotalPriceWithoutDiscount: o.State.ApplyTax(orderValue),
		TotalPrice:                o.State.ApplyTax(taxableValue),
	}
}

type Receipt struct {
	OrderValue                Dollar
	DiscountValue             Dollar
	TaxValue                  Dollar
	TotalPriceWithoutDiscount Dollar
	TotalPrice                Dollar
}
