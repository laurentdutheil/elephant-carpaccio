package controller_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/controller"
)

func TestGenerateOrder(t *testing.T) {
	t.Run("should generate an order with stateCode", func(t *testing.T) {
		orderGenerator := NewOrderGenerator(nil)

		order := orderGenerator.GenerateOrder(NoDiscount, AL)

		assert.Equal(t, AL.State(), order.State)
	})

	t.Run("should generate nbItems greater than Decimal(0)", func(t *testing.T) {
		alwaysZeroRandFunc := func(_ int64) int64 { return 0 }
		randomizer := NewDecimalRandomizer(alwaysZeroRandFunc)
		orderGenerator := NewOrderGenerator(randomizer)

		order := orderGenerator.GenerateOrder(NoDiscount, UT)

		assert.Greater(t, order.NumberOfItems, Decimal(0))
	})

	t.Run("should generate nbItems lower than Decimal(10000)", func(t *testing.T) {
		alwaysMaxRandFunc := func(n int64) int64 { return n - 1 }
		randomizer := NewDecimalRandomizer(alwaysMaxRandFunc)
		orderGenerator := NewOrderGenerator(randomizer)

		order := orderGenerator.GenerateOrder(NoDiscount, UT)

		assert.Less(t, order.NumberOfItems, Decimal(10000))
	})

	t.Run("should generate discount order greater or equal to minimal Discount amount", func(t *testing.T) {
		alwaysZeroRandFunc := func(_ int64) int64 { return 0 }
		randomizer := NewDecimalRandomizer(alwaysZeroRandFunc)
		orderGenerator := NewOrderGenerator(randomizer)
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
				order := orderGenerator.GenerateOrder(test.discountLevel, UT)

				receipt := order.Compute()
				actualOrderValue := receipt.OrderValue

				minAmount, _ := test.discountLevel.AmountRange()
				assert.True(t, actualOrderValue.GreaterOrEqual(minAmount), "%v should be greater or equal than %v", actualOrderValue, minAmount)
			})
		}
	})

	t.Run("should generate discount order lower than maximal Discount amount", func(t *testing.T) {
		alwaysMaxRandFunc := func(n int64) int64 { return n - 1 }
		randomizer := NewDecimalRandomizer(alwaysMaxRandFunc)
		orderGenerator := NewOrderGenerator(randomizer)
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
				order := orderGenerator.GenerateOrder(test.discountLevel, UT)

				receipt := order.Compute()
				actualOrderValue := receipt.OrderValue

				_, maxAmount := test.discountLevel.AmountRange()
				assert.True(t, actualOrderValue.Lower(maxAmount), "%v should be lower than %v", actualOrderValue, maxAmount)
			})
		}
	})
}

func TestPickStateCode(t *testing.T) {
	t.Run("should pick a state at random", func(t *testing.T) {
		for stateCode := 0; stateCode < int(NumberOfStates); stateCode++ {
			state := StateCode(stateCode).State()
			t.Run(state.Label, func(t *testing.T) {
				fixedIntRandom := func(_ int64) int64 { return int64(stateCode) }
				randomizer := NewDecimalRandomizer(fixedIntRandom)
				orderGenerator := NewOrderGenerator(randomizer)

				pickStateCode := orderGenerator.PickStateCode()

				assert.Equal(t, state, pickStateCode.State())
			})
		}
	})
}

func TestPickDiscountLevel(t *testing.T) {
	t.Run("should pick a discount level at random", func(t *testing.T) {
		for discountLevel := DiscountLevel(0); discountLevel < NumberOfDiscounts; discountLevel++ {
			discount := discountLevel.Discount()
			t.Run(discount.Rate.String(), func(t *testing.T) {
				fixedIntRandom := func(_ int64) int64 { return int64(discountLevel) }
				randomizer := NewDecimalRandomizer(fixedIntRandom)
				orderGenerator := NewOrderGenerator(randomizer)

				pickDiscountLevel := orderGenerator.PickDiscountLevel()

				assert.Equal(t, discount, pickDiscountLevel.Discount())
			})
		}
	})
}
