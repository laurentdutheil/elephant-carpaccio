package controller_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/controller"
)

func TestCalculator(t *testing.T) {
	t.Run("should compute order value", func(t *testing.T) {
		order := NewOrder(Decimal(1000), NewDollar(2000), UT)
		result := order.Compute()
		assert.Equal(t, NewDollar(20000), result.OrderValue)
	})

	t.Run("should compute tax for UT", func(t *testing.T) {
		order := NewOrder(Decimal(1000), NewDollar(1000), UT)
		result := order.Compute()
		assert.Equal(t, NewDollar(685), result.Tax)
	})
	t.Run("should compute tax for NV", func(t *testing.T) {
		order := NewOrder(Decimal(1000), NewDollar(1000), NV)
		result := order.Compute()
		assert.Equal(t, NewDollar(800), result.Tax)
	})
	t.Run("should compute tax for TX", func(t *testing.T) {
		order := NewOrder(Decimal(1000), NewDollar(1000), TX)
		result := order.Compute()
		assert.Equal(t, NewDollar(625), result.Tax)
	})
	t.Run("should compute tax for AL", func(t *testing.T) {
		order := NewOrder(Decimal(1000), NewDollar(1000), AL)
		result := order.Compute()
		assert.Equal(t, NewDollar(400), result.Tax)
	})
	t.Run("should compute tax for CA", func(t *testing.T) {
		order := NewOrder(Decimal(1000), NewDollar(1000), CA)
		result := order.Compute()
		assert.Equal(t, NewDollar(825), result.Tax)
	})

	t.Run("no discount for order value < $1,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(20000), UT)
		result := order.Compute()
		assert.Equal(t, NewDollar(0), result.DiscountValue)
	})
	t.Run("3% discount for order value >= $1,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(100000), UT)
		result := order.Compute()
		assert.Equal(t, NewDollar(3000), result.DiscountValue)
	})
	t.Run("5% discount for order value >= $5,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(500000), UT)
		result := order.Compute()
		assert.Equal(t, NewDollar(25000), result.DiscountValue)
	})
	t.Run("7% discount for order value >= $7,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(700000), UT)
		result := order.Compute()
		assert.Equal(t, NewDollar(49000), result.DiscountValue)
	})
	t.Run("10% discount for order value >= $10,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(1000000), UT)
		result := order.Compute()
		assert.Equal(t, NewDollar(100000), result.DiscountValue)
	})
	t.Run("15% discount for order value >= $50,000", func(t *testing.T) {
		order := NewOrder(Decimal(100), NewDollar(5000000), UT)
		result := order.Compute()
		assert.Equal(t, NewDollar(750000), result.DiscountValue)
	})

	t.Run("should compute total price", func(t *testing.T) {
		order := NewOrder(Decimal(1000), NewDollar(1000), UT)
		result := order.Compute()
		assert.Equal(t, NewDollar(10685), result.TotalPrice)
	})
	t.Run("should compute total price with discount", func(t *testing.T) {
		order := NewOrder(Decimal(10000), NewDollar(12550), AL)
		result := order.Compute()
		assert.Equal(t, NewDollar(1174680), result.TotalPrice)
	})

	t.Run("should compute taxed price without discount", func(t *testing.T) {
		order := NewOrder(Decimal(10000), NewDollar(12550), AL)
		result := order.Compute()
		assert.Equal(t, NewDollar(1305200), result.TaxedOrderValue)
	})
}
