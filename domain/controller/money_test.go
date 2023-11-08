package controller_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/controller"
)

func TestDollar_String_should_add_dollar_at_the_beginning_of_decimal_String(t *testing.T) {
	decimal := Decimal(123)
	dollar := NewDollar(decimal)
	assert.Equal(t, "$"+decimal.String(), dollar.String())
}

func TestDecimal(t *testing.T) {
	t.Run("should Multiply Two Decimals", func(t *testing.T) {
		first := Decimal(100)
		second := Decimal(300)

		assert.Equal(t, Decimal(300), first.Multiply(second))
	})
	t.Run("should Multiply Two Decimals with rounding", func(t *testing.T) {
		first := Decimal(123)
		second := Decimal(321)

		assert.Equal(t, Decimal(395), first.Multiply(second))
	})
	t.Run("should Divide Two Decimals", func(t *testing.T) {
		first := Decimal(900)
		second := Decimal(300)

		assert.Equal(t, Decimal(300), first.Divide(second))
	})
	t.Run("should Divide Two Decimals with rounding", func(t *testing.T) {
		first := Decimal(321)
		second := Decimal(123)

		assert.Equal(t, Decimal(261), first.Divide(second))
	})
}
