package controller_test

import (
	. "elephant_carpaccio/domain/money"
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/controller"
)

func TestGenerateOrder(t *testing.T) {
	t.Run("should generate an order with with", func(t *testing.T) {
		orderGenerator := NewOrderGenerator(nil)

		order := orderGenerator.GenerateOrder(No.Discount(), AL.State())

		assert.Equal(t, AL.State(), order.State)
	})

	t.Run("should generate nbItems greater than Decimal(0)", func(t *testing.T) {
		alwaysZeroRandFunc := func(_ int64) int64 { return 0 }
		randomizer := NewDecimalRandomizer(alwaysZeroRandFunc)
		orderGenerator := NewOrderGenerator(randomizer)

		order := orderGenerator.GenerateOrder(No.Discount(), UT.State())

		assert.Greater(t, order.NumberOfItems, Decimal(0))
	})

	t.Run("should generate nbItems lower than Decimal(10000)", func(t *testing.T) {
		alwaysMaxRandFunc := func(n int64) int64 { return n - 1 }
		randomizer := NewDecimalRandomizer(alwaysMaxRandFunc)
		orderGenerator := NewOrderGenerator(randomizer)

		order := orderGenerator.GenerateOrder(No.Discount(), UT.State())

		assert.Less(t, order.NumberOfItems, Decimal(10000))
	})

	t.Run("should generate discount order greater or equal to minimal Discount amount", func(t *testing.T) {
		alwaysZeroRandFunc := func(_ int64) int64 { return 0 }
		randomizer := NewDecimalRandomizer(alwaysZeroRandFunc)
		orderGenerator := NewOrderGenerator(randomizer)
		tests := []struct {
			description string
			discount    *Discount
		}{
			{"should generate a no discount order", No.Discount()},
			{"should generate a 3% discount order", ThreePercent.Discount()},
			{"should generate a 5% discount order", FivePercent.Discount()},
			{"should generate a 7% discount order", SevenPercent.Discount()},
			{"should generate a 10% discount order", TenPercent.Discount()},
			{"should generate a 15% discount order", FifteenPercent.Discount()},
		}
		for _, test := range tests {
			t.Run(test.description, func(t *testing.T) {
				order := orderGenerator.GenerateOrder(test.discount, UT.State())

				receipt := order.Compute()
				actualOrderValue := receipt.OrderValue

				minAmount, _ := test.discount.AmountRange()
				assert.True(t, actualOrderValue.GreaterOrEqual(minAmount), "%v should be greater or equal than %v", actualOrderValue, minAmount)
			})
		}
	})

	t.Run("should generate discount order lower than maximal Discount amount", func(t *testing.T) {
		alwaysMaxRandFunc := func(n int64) int64 { return n - 1 }
		randomizer := NewDecimalRandomizer(alwaysMaxRandFunc)
		orderGenerator := NewOrderGenerator(randomizer)
		tests := []struct {
			description string
			discount    *Discount
		}{
			{"should generate a no discount order", No.Discount()},
			{"should generate a 3% discount order", ThreePercent.Discount()},
			{"should generate a 5% discount order", FivePercent.Discount()},
			{"should generate a 7% discount order", SevenPercent.Discount()},
			{"should generate a 10% discount order", TenPercent.Discount()},
			{"should generate a 15% discount order", FifteenPercent.Discount()},
		}
		for _, test := range tests {
			t.Run(test.description, func(t *testing.T) {
				order := orderGenerator.GenerateOrder(test.discount, UT.State())

				receipt := order.Compute()
				actualOrderValue := receipt.OrderValue

				_, maxAmount := test.discount.AmountRange()
				assert.True(t, actualOrderValue.Lower(maxAmount), "%v should be lower than %v", actualOrderValue, maxAmount)
			})
		}
	})

	t.Run("should pick a state at random when argument is nil", func(t *testing.T) {
		for i, state := range AllStates() {
			t.Run(state.Label, func(t *testing.T) {
				fixedIntRandom := func(_ int64) int64 { return int64(i) }
				randomizer := NewDecimalRandomizer(fixedIntRandom)
				orderGenerator := NewOrderGenerator(randomizer)

				order := orderGenerator.GenerateOrder(No.Discount(), nil)

				assert.Equal(t, &state, order.State)
			})
		}
	})

	t.Run("should pick a discount level at random when argument is nil", func(t *testing.T) {
		for i, discount := range AllDiscounts() {
			t.Run(discount.Rate.String(), func(t *testing.T) {
				fixedIntRandom := func(_ int64) int64 { return int64(i) }
				randomizer := NewDecimalRandomizer(fixedIntRandom)
				orderGenerator := NewOrderGenerator(randomizer)

				order := orderGenerator.GenerateOrder(nil, AL.State())
				receipt := order.Compute()

				assert.Equal(t, discount, receipt.Discount)
			})
		}
	})

}
