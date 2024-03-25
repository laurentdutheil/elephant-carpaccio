package money_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/money"
)

func TestPercent_String_should_add_percent_at_the_end_of_decimal_String(t *testing.T) {
	decimal := Decimal(123)
	percent := NewPercent(decimal)
	assert.Equal(t, decimal.String()+"%", percent.String())
}

func TestPercent_ApplyTo(t *testing.T) {
	percent := NewPercent(Decimal(1000))
	amount := NewDollar(Decimal(20000))

	discount := percent.ApplyTo(amount)

	assert.Equal(t, NewDollar(Decimal(2000)), discount)
}
