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
	discount := ComputeDiscount(orderValue)
	taxableValue := orderValue.Minus(discount)
	return Receipt{
		OrderValue:      orderValue,
		Discount:        discount,
		Tax:             o.State.ComputeTax(taxableValue),
		TaxedOrderValue: o.State.ApplyTax(orderValue),
		TotalPrice:      o.State.ApplyTax(taxableValue),
	}
}

type Receipt struct {
	OrderValue      Dollar
	Discount        Dollar
	Tax             Dollar
	TaxedOrderValue Dollar
	TotalPrice      Dollar
}
