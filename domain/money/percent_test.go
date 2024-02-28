package money_test

import (
	"elephant_carpaccio/domain/money"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPercent_String_should_add_percent_at_the_end_of_decimal_String(t *testing.T) {
	decimal := money.Decimal(123)
	percent := money.NewPercent(decimal)
	assert.Equal(t, decimal.String()+"%", percent.String())
}

func TestPercent_ApplyTo(t *testing.T) {
	percent := money.NewPercent(money.Decimal(1000))
	amount := money.NewDollar(money.Decimal(20000))

	discount := percent.ApplyTo(amount)

	assert.Equal(t, money.NewDollar(money.Decimal(2000)), discount)
}
