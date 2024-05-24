package calculator_test

import (
	"elephant_carpaccio/domain/money"
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain/calculator"
)

func TestState(t *testing.T) {
	tests := []struct {
		name          string
		stateCode     StateCode
		expectedLabel string
		expectedTax   money.Percent
	}{
		{"Utah have 'UT' label and 6.85% tax", UT, "UT", money.NewPercent(685)},
		{"Nevada have 'NV' label and 8.00% tax", NV, "NV", money.NewPercent(800)},
		{"Texas have 'TX' label and 6.25% tax", TX, "TX", money.NewPercent(625)},
		{"Alabama have 'AL' label and 4.00% tax", AL, "AL", money.NewPercent(400)},
		{"California have 'CA' label and 8.25% tax", CA, "CA", money.NewPercent(825)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			state := StateOf(test.stateCode)
			assert.Equal(t, test.expectedLabel, state.Label)
			assert.Equal(t, test.expectedTax, state.TaxRate)
		})
	}

	t.Run("StateOf should return nil if code does not exist", func(t *testing.T) {
		assert.Nil(t, StateOf(StateCode(123)))
	})
}

func TestState_ComputeTax(t *testing.T) {
	tests := []struct {
		name                string
		stateCode           StateCode
		expectedComputedTax money.Dollar
	}{
		{"Utah compute a 6.85% tax for a $100.00 amount", UT, money.NewDollar(685)},
		{"Nevada compute a 8.00% tax for a $100.00 amount", NV, money.NewDollar(800)},
		{"Texas compute a 6.25% tax for a $100.00 amount", TX, money.NewDollar(625)},
		{"Alabama compute a 4.00% tax for a $100.00 amount", AL, money.NewDollar(400)},
		{"California compute a 8.25% tax for a $100.00 amount", CA, money.NewDollar(825)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			state := StateOf(test.stateCode)
			assert.Equal(t, test.expectedComputedTax, state.ComputeTax(money.NewDollar(10000)))
		})
	}
}

func TestState_ApplyTax(t *testing.T) {
	tests := []struct {
		name                string
		stateCode           StateCode
		expectedTaxedAmount money.Dollar
	}{
		{"Utah apply a 6.85% tax for a $100.00 amount", UT, money.NewDollar(10685)},
		{"Nevada apply a 8.00% tax for a $100.00 amount", NV, money.NewDollar(10800)},
		{"Texas apply a 6.25% tax for a $100.00 amount", TX, money.NewDollar(10625)},
		{"Alabama apply a 4.00% tax for a $100.00 amount", AL, money.NewDollar(10400)},
		{"California apply a 8.25% tax for a $100.00 amount", CA, money.NewDollar(10825)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			state := StateOf(test.stateCode)
			assert.Equal(t, test.expectedTaxedAmount, state.ApplyTax(money.NewDollar(10000)))
		})
	}
}
