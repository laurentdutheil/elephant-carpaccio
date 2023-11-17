package controller

type Order struct {
	NumberOfItems Decimal
	ItemPrice     Dollar
	State         *State
}

func NewOrder(numberOfItems Decimal, itemPrice Dollar, state *State) Order {
	return Order{NumberOfItems: numberOfItems, ItemPrice: itemPrice, State: state}
}

func (o Order) Compute() Receipt {
	orderValue := o.ItemPrice.Multiply(o.NumberOfItems)
	discountValue, discount := ComputeDiscount(orderValue)
	taxableValue := orderValue.Minus(discountValue)
	return Receipt{
		OrderValue:      orderValue,
		Discount:        discount,
		DiscountValue:   discountValue,
		Tax:             o.State.ComputeTax(taxableValue),
		TaxedOrderValue: o.State.ApplyTax(orderValue),
		TotalPrice:      o.State.ApplyTax(taxableValue),
	}
}

type Receipt struct {
	OrderValue      Dollar
	Discount        Discount
	DiscountValue   Dollar
	Tax             Dollar
	TaxedOrderValue Dollar
	TotalPrice      Dollar
}
