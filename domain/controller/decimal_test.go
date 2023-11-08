package controller_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/controller"
)

func TestDecimal_String(t *testing.T) {
	tests := []struct {
		name     string
		decimal  Decimal
		expected string
	}{
		{"should put the separator for two decimals number", Decimal(123), "1.23"},
		{"should add a zero if decimals have one digit", Decimal(103), "1.03"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.decimal.String(), fmt.Sprint(test.decimal))
		})
	}
}

func TestPercent_String_should_add_percent_at_the_end_of_decimal_String(t *testing.T) {
	decimal := Decimal(123)
	percent := NewPercent(decimal)
	assert.Equal(t, decimal.String()+"%", percent.String())
}
