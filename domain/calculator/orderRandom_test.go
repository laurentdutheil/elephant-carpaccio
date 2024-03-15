package calculator_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/calculator"
	. "elephant_carpaccio/domain/money"
)

func TestOrderRandomizer_RandDecimal(t *testing.T) {
	t.Run("should generate decimal greater or equal than min parameter", func(t *testing.T) {
		randomizer := NewOrderRandomizer()
		RandInt63n = func(_ int64) int64 { return 0 }

		randDecimal := randomizer.RandDecimal(Decimal(12), Decimal(20))

		assert.Equal(t, randDecimal, Decimal(12))
	})

	t.Run("should generate decimal lower than max parameter", func(t *testing.T) {
		randomizer := NewOrderRandomizer()
		RandInt63n = func(n int64) int64 { return n - 1 }

		randDecimal := randomizer.RandDecimal(Decimal(12), Decimal(20))

		assert.Equal(t, randDecimal, Decimal(19))
	})
}

func TestOrderRandomizer_RandDecimalWithoutDecimals(t *testing.T) {
	t.Run("should generate decimal greater or equal than min parameter", func(t *testing.T) {
		randomizer := NewOrderRandomizer()
		RandInt63n = func(_ int64) int64 { return 0 }

		randDecimal := randomizer.RandDecimalWithoutDecimals(Decimal(123), Decimal(259))

		assert.Equal(t, randDecimal, Decimal(100))
	})

	t.Run("should generate decimal lower than max parameter", func(t *testing.T) {
		randomizer := NewOrderRandomizer()
		RandInt63n = func(n int64) int64 { return n - 1 }

		randDecimal := randomizer.RandDecimalWithoutDecimals(Decimal(123), Decimal(259))

		assert.Equal(t, randDecimal, Decimal(200))
	})
}

func TestOrderRandomizer_RandDollar(t *testing.T) {
	t.Run("should generate dollar greater or equal than minAmount parameter", func(t *testing.T) {
		randomizer := NewOrderRandomizer()
		RandInt63n = func(_ int64) int64 { return 0 }

		randDollar := randomizer.RandDollar(NewDollar(Decimal(1200)), NewDollar(Decimal(2000)))

		assert.Equal(t, randDollar, NewDollar(Decimal(1200)))
	})

	t.Run("should generate dollar lower than max parameter", func(t *testing.T) {
		randomizer := NewOrderRandomizer()
		RandInt63n = func(n int64) int64 { return n - 1 }

		randDollar := randomizer.RandDollar(NewDollar(Decimal(1200)), NewDollar(Decimal(2000)))

		assert.Equal(t, randDollar, NewDollar(Decimal(1999)))
	})
}

func TestOrderRandomizer_RandState(t *testing.T) {
	tests := []struct {
		name      string
		stateCode StateCode
	}{
		{"RandState return UT", UT},
		{"RandState return NV", NV},
		{"RandState return TX", TX},
		{"RandState return AL", AL},
		{"RandState return CA", CA},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			randomizer := NewOrderRandomizer()
			fixedIntRandom := func(_ int) int { return int(test.stateCode) }
			RandIntn = fixedIntRandom

			randState := randomizer.RandState()

			assert.Equal(t, test.stateCode.State(), randState)
		})
	}
}

func TestOrderRandomizer_RandDiscount(t *testing.T) {
	tests := []struct {
		name          string
		discountLevel DiscountLevel
	}{
		{"RandDiscount return No Discount", No},
		{"RandDiscount return 3% Discount", ThreePercent},
		{"RandDiscount return 5% Discount", FivePercent},
		{"RandDiscount return 7% Discount", SevenPercent},
		{"RandDiscount return 10% Discount", TenPercent},
		{"RandDiscount return 15% Discount", FifteenPercent},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			randomizer := NewOrderRandomizer()
			fixedIntRandom := func(_ int) int { return int(test.discountLevel) }
			RandIntn = fixedIntRandom

			randDiscountLevel := randomizer.RandDiscountLevel()

			assert.Equal(t, test.discountLevel.Discount(), randDiscountLevel)
		})
	}
}
