package money_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/money"
)

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

func TestDecimal_String(t *testing.T) {
	tests := []struct {
		name     string
		decimal  Decimal
		expected string
	}{
		{"should put the decimal separator for two decimals number", Decimal(123), "1.23"},
		{"should add a zero if decimals have one digit", Decimal(103), "1.03"},
		{"should put the thousand separator", Decimal(165400023), "1,654,000.23"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, fmt.Sprint(test.decimal))
		})
	}
}

func TestDecimal_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		decimal  Decimal
		expected string
	}{
		{"should put the decimal separator for two decimals number", Decimal(123), "1.23"},
		{"should add a zero if decimals have one digit", Decimal(103), "1.03"},
		{"should not put the thousand separator", Decimal(165400023), "1654000.23"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			json, _ := test.decimal.MarshalJSON()
			assert.Equal(t, test.expected, fmt.Sprint(string(json)))
		})
	}
}
