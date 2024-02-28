package money_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/money"
)

func TestDollar_String_should_add_dollar_at_the_beginning_of_decimal_String(t *testing.T) {
	dollar := NewDollar(Decimal(123))
	assert.Equal(t, "$"+dollar.AmountInCents().String(), dollar.String())
}

func TestDollar_GreaterOrEqual(t *testing.T) {
	tests := []struct {
		name     string
		actual   Dollar
		other    Dollar
		expected bool
	}{
		{"lower amount should return false", NewDollar(Decimal(1999)), NewDollar(Decimal(2000)), false},
		{"equal amount should return true", NewDollar(Decimal(2000)), NewDollar(Decimal(2000)), true},
		{"greater amount should return true", NewDollar(Decimal(2001)), NewDollar(Decimal(2000)), true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			greaterOrEqual := test.actual.GreaterOrEqual(test.other)
			assert.Equal(t, test.expected, greaterOrEqual)
		})
	}
}
func TestDollar_Lower(t *testing.T) {
	tests := []struct {
		name     string
		actual   Dollar
		other    Dollar
		expected bool
	}{
		{"lower amount should return true", NewDollar(Decimal(1999)), NewDollar(Decimal(2000)), true},
		{"equal amount should return false", NewDollar(Decimal(2000)), NewDollar(Decimal(2000)), false},
		{"greater amount should return false", NewDollar(Decimal(2001)), NewDollar(Decimal(2000)), false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			greaterOrEqual := test.actual.Lower(test.other)
			assert.Equal(t, test.expected, greaterOrEqual)
		})
	}
}

func TestDollar_Add(t *testing.T) {
	sum := NewDollar(Decimal(2000)).Add(NewDollar(Decimal(1000)))
	assert.Equal(t, NewDollar(Decimal(3000)), sum)
}

func TestDollar_Minus(t *testing.T) {
	diff := NewDollar(Decimal(2000)).Minus(NewDollar(Decimal(500)))
	assert.Equal(t, NewDollar(Decimal(1500)), diff)
}
