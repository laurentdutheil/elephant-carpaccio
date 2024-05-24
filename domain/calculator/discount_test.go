package calculator_test

import (
	"elephant_carpaccio/domain/money"
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/calculator"
)

func TestDiscount(t *testing.T) {
	tests := []struct {
		name                 string
		discountLevel        DiscountLevel
		expectedMinAmount    money.Dollar
		expectedMaxAmount    money.Dollar
		expectedDiscountRate money.Percent
	}{
		{"no discount for order value < $1,000", No, money.NewDollar(0), money.NewDollar(100000), money.NewPercent(0)},
		{"3% discount for order value >= $1,000", ThreePercent, money.NewDollar(100000), money.NewDollar(500000), money.NewPercent(300)},
		{"5% discount for order value >= $5,000", FivePercent, money.NewDollar(500000), money.NewDollar(700000), money.NewPercent(500)},
		{"7% discount for order value >= $7,000", SevenPercent, money.NewDollar(700000), money.NewDollar(1000000), money.NewPercent(700)},
		{"10% discount for order value >= $10,000", TenPercent, money.NewDollar(1000000), money.NewDollar(5000000), money.NewPercent(1000)},
		{"15% discount for order value >= $50,000", FifteenPercent, money.NewDollar(5000000), money.NewDollar(100000000), money.NewPercent(1500)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			discount := DiscountOf(test.discountLevel)
			minAmount, maxAmount := discount.AmountRange()
			assert.Equal(t, test.expectedMinAmount, minAmount)
			assert.Equal(t, test.expectedMaxAmount, maxAmount)
			assert.Equal(t, test.expectedDiscountRate, discount.Rate)
		})
	}

	t.Run("DiscountOf should return nil if code does not exist", func(t *testing.T) {
		assert.Nil(t, DiscountOf(DiscountLevel(123)))
	})
}

func TestDiscount_ComputeDiscount(t *testing.T) {
	tests := []struct {
		name                     string
		discountLevel            DiscountLevel
		expectedComputedDiscount money.Dollar
	}{
		{"compute no discount for a $100.00 amount", No, money.NewDollar(0)},
		{"compute 3% discount for a $100.00 amount", ThreePercent, money.NewDollar(300)},
		{"compute 5% discount for a $100.00 amount", FivePercent, money.NewDollar(500)},
		{"compute 7% discount for a $100.00 amount", SevenPercent, money.NewDollar(700)},
		{"compute 10% discount for a $100.00 amount", TenPercent, money.NewDollar(1000)},
		{"compute 15% discount for a $100.00 amount", FifteenPercent, money.NewDollar(1500)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			discount := DiscountOf(test.discountLevel)
			assert.Equal(t, test.expectedComputedDiscount, discount.ComputeDiscount(money.NewDollar(10000)))
		})
	}
}

func TestFindDiscount(t *testing.T) {
	tests := []struct {
		name                  string
		amount                money.Dollar
		expectedDiscountLevel DiscountLevel
	}{
		{"no discount for order value < $1,000", money.NewDollar(0), No},
		{"no discount for order value < $1,000", money.NewDollar(99999), No},
		{"3% discount for order value >= $1,000", money.NewDollar(100000), ThreePercent},
		{"3% discount for order value >= $1,000", money.NewDollar(499999), ThreePercent},
		{"5% discount for order value >= $5,000", money.NewDollar(500000), FivePercent},
		{"5% discount for order value >= $5,000", money.NewDollar(699999), FivePercent},
		{"7% discount for order value >= $7,000", money.NewDollar(700000), SevenPercent},
		{"7% discount for order value >= $7,000", money.NewDollar(999999), SevenPercent},
		{"10% discount for order value >= $10,000", money.NewDollar(1000000), TenPercent},
		{"10% discount for order value >= $10,000", money.NewDollar(4999999), TenPercent},
		{"15% discount for order value >= $50,000", money.NewDollar(5000000), FifteenPercent},
		{"15% discount for order value >= $50,000", money.NewDollar(99999999), FifteenPercent},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			discount := FindDiscount(test.amount)
			assert.Equal(t, DiscountOf(test.expectedDiscountLevel), discount)
		})
	}

	t.Run("DiscountOf should return nil if code does not exist", func(t *testing.T) {
		assert.Nil(t, DiscountOf(DiscountLevel(123)))
	})
}
