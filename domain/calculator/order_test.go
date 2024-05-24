package calculator_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/calculator"
	. "elephant_carpaccio/domain/money"
)

func TestOrder(t *testing.T) {
	t.Run("should compute order value", func(t *testing.T) {
		order := NewOrder(Decimal(1000), NewDollar(2000), UT.State())

		result := order.Compute()

		assert.Equal(t, NewDollar(20000), result.OrderValue)
	})

	t.Run("should compute total price", func(t *testing.T) {
		order := NewOrder(Decimal(1000), NewDollar(1000), UT.State())

		result := order.Compute()

		assert.Equal(t, NewDollar(10685), result.TotalPrice)
	})

	t.Run("should compute total price with discount", func(t *testing.T) {
		order := NewOrder(Decimal(10000), NewDollar(12550), AL.State())

		result := order.Compute()

		assert.Equal(t, NewDollar(1174680), result.TotalPrice)
	})

	t.Run("should compute taxed price without discount", func(t *testing.T) {
		order := NewOrder(Decimal(10000), NewDollar(12550), AL.State())

		result := order.Compute()

		assert.Equal(t, NewDollar(1305200), result.TaxedOrderValue)
	})
}

func TestOrder_Compute_Tax(t *testing.T) {
	tests := []struct {
		name        string
		stateCode   StateCode
		expectedTax Dollar
	}{
		{"should compute tax for ", UT, NewDollar(685)},
		{"should compute tax for ", NV, NewDollar(800)},
		{"should compute tax for ", TX, NewDollar(625)},
		{"should compute tax for ", AL, NewDollar(400)},
		{"should compute tax for ", CA, NewDollar(825)},
	}
	for _, test := range tests {
		t.Run(test.name+test.stateCode.State().Label, func(t *testing.T) {
			order := NewOrder(Decimal(1000), NewDollar(1000), test.stateCode.State())

			result := order.Compute()

			assert.Equal(t, test.expectedTax, result.Tax)
			assert.Equal(t, NewDollar(10000).Add(test.expectedTax), result.TaxedOrderValue)
		})
	}
}

func TestOrder_Compute_Discount(t *testing.T) {
	t.Run("no discount for order value < $1,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(20000), UT.State())

		result := order.Compute()

		assert.Equal(t, No.Discount(), result.Discount)
		assert.Equal(t, NewDollar(0), result.DiscountValue)
	})
	t.Run("3% discount for order value >= $1,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(100000), UT.State())

		result := order.Compute()

		assert.Equal(t, ThreePercent.Discount(), result.Discount)
		assert.Equal(t, NewDollar(3000), result.DiscountValue)
	})
	t.Run("5% discount for order value >= $5,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(500000), UT.State())

		result := order.Compute()

		assert.Equal(t, FivePercent.Discount(), result.Discount)
		assert.Equal(t, NewDollar(25000), result.DiscountValue)
	})
	t.Run("7% discount for order value >= $7,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(700000), UT.State())

		result := order.Compute()

		assert.Equal(t, SevenPercent.Discount(), result.Discount)
		assert.Equal(t, NewDollar(49000), result.DiscountValue)
	})
	t.Run("10% discount for order value >= $10,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(1000000), UT.State())

		result := order.Compute()

		assert.Equal(t, TenPercent.Discount(), result.Discount)
		assert.Equal(t, NewDollar(100000), result.DiscountValue)
	})
	t.Run("15% discount for order value >= $50,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(5000000), UT.State())

		result := order.Compute()

		assert.Equal(t, FifteenPercent.Discount(), result.Discount)
		assert.Equal(t, NewDollar(750000), result.DiscountValue)
	})
}
