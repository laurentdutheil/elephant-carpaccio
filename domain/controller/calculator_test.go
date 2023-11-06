package controller_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"

	. "elephant_carpaccio/domain/controller"
)

func TestCalculator(t *testing.T) {
	t.Run("should compute order value", func(t *testing.T) {
		result := Compute(10.0, NewDollar(2000), "UT")
		assert.Equal(t, NewDollar(20000), result.OrderValue)
	})

	t.Run("should compute tax for UT", func(t *testing.T) {
		result := Compute(10.0, NewDollar(1000), "UT")
		assert.Equal(t, NewDollar(685), result.TaxValue)
	})
	t.Run("should compute tax for NV", func(t *testing.T) {
		result := Compute(10.0, NewDollar(1000), "NV")
		assert.Equal(t, NewDollar(800), result.TaxValue)
	})
	t.Run("should compute tax for TX", func(t *testing.T) {
		result := Compute(10.0, NewDollar(1000), "TX")
		assert.Equal(t, NewDollar(625), result.TaxValue)
	})
	t.Run("should compute tax for AL", func(t *testing.T) {
		result := Compute(10.0, NewDollar(1000), "AL")
		assert.Equal(t, NewDollar(400), result.TaxValue)
	})
	t.Run("should compute tax for CA", func(t *testing.T) {
		result := Compute(10.0, NewDollar(1000), "CA")
		assert.Equal(t, NewDollar(825), result.TaxValue)
	})

	t.Run("no discount for order value < $1,000", func(t *testing.T) {
		result := Compute(1.0, NewDollar(20000), "UT")
		assert.Equal(t, NewDollar(0), result.DiscountValue)
	})
	t.Run("3% discount for order value >= $1,000", func(t *testing.T) {
		result := Compute(1.0, NewDollar(100000), "UT")
		assert.Equal(t, NewDollar(3000), result.DiscountValue)
	})
	t.Run("5% discount for order value >= $5,000", func(t *testing.T) {
		result := Compute(1.0, NewDollar(500000), "UT")
		assert.Equal(t, NewDollar(25000), result.DiscountValue)
	})
	t.Run("7% discount for order value >= $7,000", func(t *testing.T) {
		result := Compute(1.0, NewDollar(700000), "UT")
		assert.Equal(t, NewDollar(49000), result.DiscountValue)
	})
	t.Run("10% discount for order value >= $10,000", func(t *testing.T) {
		result := Compute(1.0, NewDollar(1000000), "UT")
		assert.Equal(t, NewDollar(100000), result.DiscountValue)
	})
	t.Run("15% discount for order value >= $50,000", func(t *testing.T) {
		result := Compute(1.0, NewDollar(5000000), "UT")
		assert.Equal(t, NewDollar(750000), result.DiscountValue)
	})

	t.Run("should compute total price", func(t *testing.T) {
		result := Compute(10.0, NewDollar(1000), "UT")
		assert.Equal(t, NewDollar(10685), result.TotalPrice)
	})
	t.Run("should compute total price with discount", func(t *testing.T) {
		result := Compute(100.0, NewDollar(12550), "AL")
		assert.Equal(t, NewDollar(1174680), result.TotalPrice)
	})

	t.Run("should compute taxed price without discount", func(t *testing.T) {
		result := Compute(100.0, NewDollar(12550), "AL")
		assert.Equal(t, NewDollar(1305200), result.TotalPriceWithoutDiscount)
	})
}

func TestGenerateOrder(t *testing.T) {
	t.Run("should generate nbItems with only two decimals", func(t *testing.T) {
		nbItems, _ := GenerateOrder(NoDiscount)
		assert.Equal(t, math.Trunc(nbItems*math.Pow10(2)), nbItems*math.Pow10(2), "%v should have only two decimals", nbItems)
	})

	t.Run("should generate all discount order", func(t *testing.T) {
		tests := []struct {
			description   string
			discountLevel DiscountLevel
		}{
			{"should generate a no discount order", NoDiscount},
			{"should generate a 3% discount order", ThreePercentDiscount},
			{"should generate a 5% discount order", FivePercentDiscount},
			{"should generate a 7% discount order", SevenPercentDiscount},
			{"should generate a 10% discount order", TenPercentDiscount},
			{"should generate a 15% discount order", FifteenPercentDiscount},
		}
		for _, test := range tests {
			t.Run(test.description, func(t *testing.T) {
				nbItems, itemPrice := GenerateOrder(test.discountLevel)
				result := itemPrice.Multiply(nbItems)
				assertInDiscountLevelRange(t, result, test.discountLevel)
			})
		}
	})
}

func assertInDiscountLevelRange(t *testing.T, result Dollar, discountLevel DiscountLevel) {
	minAmount, maxAmount := discountLevel.AmountRange()
	assert.True(t, result.GreaterOrEqual(minAmount), "%v should be greater or equal than %v", result, minAmount)
	assert.True(t, result.Lower(maxAmount), "%v should be lower than %v", result, maxAmount)
}
