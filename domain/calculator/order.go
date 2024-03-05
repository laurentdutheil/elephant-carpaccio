package calculator

import "elephant_carpaccio/domain/money"

type Order struct {
	NumberOfItems money.Decimal
	ItemPrice     money.Dollar
	State         *State
}

func NewOrder(numberOfItems money.Decimal, itemPrice money.Dollar, state *State) Order {
	return Order{NumberOfItems: numberOfItems, ItemPrice: itemPrice, State: state}
}

func (o Order) Compute() Receipt {
	orderValue := o.ItemPrice.Multiply(o.NumberOfItems)
	discount := FindDiscount(orderValue)
	discountValue := discount.ComputeDiscount(orderValue)
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
	OrderValue      money.Dollar
	Discount        *Discount
	DiscountValue   money.Dollar
	Tax             money.Dollar
	TaxedOrderValue money.Dollar
	TotalPrice      money.Dollar
}
