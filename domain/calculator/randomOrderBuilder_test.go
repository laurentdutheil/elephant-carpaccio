package calculator_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"regexp"
	"runtime"
	"strings"
	"testing"

	. "elephant_carpaccio/domain/calculator"
	. "elephant_carpaccio/domain/money"
)

func TestBuildOrder(t *testing.T) {
	t.Run("should build an order with fixed Discount and fixed State", func(t *testing.T) {
		orderBuilder := NewRandomOrderBuilder(NewOrderRandomizer()).
			WithDiscount(No.Discount()).
			WithState(AL.State())

		order := orderBuilder.Build()
		receipt := order.Compute()

		assert.Equal(t, AL.State(), order.State)
		assert.Equal(t, No.Discount(), receipt.Discount)
	})

	t.Run("should build a ramdom nbItems between Decimal(1) included and Decimal(10000) excluded", func(t *testing.T) {
		spyOrderRandom := NewSpyOrderRandom(NewOrderRandomizer())
		orderBuilder := NewRandomOrderBuilder(spyOrderRandom).
			WithDiscount(No.Discount()).
			WithState(UT.State())

		order := orderBuilder.Build()

		assert.GreaterOrEqual(t, order.NumberOfItems, Decimal(1))
		assert.Less(t, order.NumberOfItems, Decimal(10000))
		spyOrderRandom.AssertCalled(t, "RandDecimal", Decimal(1), Decimal(10000))
	})

	t.Run("should build a ramdom nbItems without decimals between Decimal(100) included and Decimal(10000) excluded", func(t *testing.T) {
		spyOrderRandom := NewSpyOrderRandom(NewOrderRandomizer())
		orderBuilder := NewRandomOrderBuilder(spyOrderRandom).
			WithDiscount(No.Discount()).
			WithState(UT.State()).
			WithoutDecimals(true)

		order := orderBuilder.Build()

		assert.GreaterOrEqual(t, order.NumberOfItems, Decimal(100))
		assert.Less(t, order.NumberOfItems, Decimal(10000))
		assert.True(t, strings.HasSuffix(order.NumberOfItems.String(), ".00"), "the number of items should not have decimals")
		spyOrderRandom.AssertCalled(t, "RandDecimalWithoutDecimals", Decimal(1), Decimal(10000))
	})

	t.Run("should build a random discount order between minimal Discount amount and maximal Discount amount", func(t *testing.T) {
		spyOrderRandom := NewSpyOrderRandom(NewOrderRandomizer())
		tests := []struct {
			description string
			discount    *Discount
		}{
			{"should build a no discount order", No.Discount()},
			{"should build a 3% discount order", ThreePercent.Discount()},
			{"should build a 5% discount order", FivePercent.Discount()},
			{"should build a 7% discount order", SevenPercent.Discount()},
			{"should build a 10% discount order", TenPercent.Discount()},
			{"should build a 15% discount order", FifteenPercent.Discount()},
		}
		for _, test := range tests {
			t.Run(test.description, func(t *testing.T) {
				orderBuilder := NewRandomOrderBuilder(spyOrderRandom).
					WithDiscount(test.discount).
					WithState(UT.State())

				order := orderBuilder.Build()
				receipt := order.Compute()
				actualOrderValue := receipt.OrderValue

				minAmount, maxAmount := test.discount.AmountRange()
				assert.True(t, actualOrderValue.GreaterOrEqual(minAmount), "%v should be greater or equal than %v", actualOrderValue, minAmount)
				assert.True(t, actualOrderValue.Lower(maxAmount), "%v should be lower than %v", actualOrderValue, maxAmount)

				minItemPrice := minAmount.Divide(order.NumberOfItems)
				maxItemPrice := maxAmount.Divide(order.NumberOfItems)
				spyOrderRandom.AssertCalled(t, "RandDollar", minItemPrice, maxItemPrice)
			})
		}
	})

	t.Run("should pick a state at random when argument is nil", func(t *testing.T) {
		spyOrderRandom := NewSpyOrderRandom(NewOrderRandomizer())
		orderBuilder := NewRandomOrderBuilder(spyOrderRandom).
			WithDiscount(No.Discount())

		order := orderBuilder.Build()
		receipt := order.Compute()

		assert.Equal(t, No.Discount(), receipt.Discount)
		assert.NotNil(t, order.State)
		spyOrderRandom.AssertCalled(t, "RandState")
	})

	t.Run("should pick a discount level at random when argument is nil", func(t *testing.T) {
		spyOrderRandom := NewSpyOrderRandom(NewOrderRandomizer())
		orderBuilder := NewRandomOrderBuilder(spyOrderRandom).
			WithState(AL.State())

		order := orderBuilder.Build()
		receipt := order.Compute()

		assert.Equal(t, AL.State(), order.State)
		assert.NotNil(t, receipt.Discount)
		spyOrderRandom.AssertCalled(t, "RandDiscountLevel")
	})

}

type SpyOrderRandom struct {
	mock.Mock
	spied OrderRandom
}

func NewSpyOrderRandom(spied OrderRandom) *SpyOrderRandom {
	return &SpyOrderRandom{spied: spied}
}

func (m *SpyOrderRandom) RandDecimal(min Decimal, max Decimal) Decimal {
	m.Calls = append(m.Calls, *NewCall(&m.Mock, min, max))
	return m.spied.RandDecimal(min, max)
}

func (m *SpyOrderRandom) RandDecimalWithoutDecimals(min Decimal, max Decimal) Decimal {
	m.Calls = append(m.Calls, *NewCall(&m.Mock, min, max))
	return m.spied.RandDecimalWithoutDecimals(min, max)
}

func (m *SpyOrderRandom) RandDollar(minAmount Dollar, maxAmount Dollar) Dollar {
	m.Calls = append(m.Calls, *NewCall(&m.Mock, minAmount, maxAmount))
	return m.spied.RandDollar(minAmount, maxAmount)
}

func (m *SpyOrderRandom) RandDiscountLevel() *Discount {
	m.Calls = append(m.Calls, *NewCall(&m.Mock))
	return m.spied.RandDiscountLevel()
}

func (m *SpyOrderRandom) RandState() *State {
	m.Calls = append(m.Calls, *NewCall(&m.Mock))
	return m.spied.RandState()
}

func NewCall(parent *mock.Mock, methodArguments ...interface{}) *mock.Call {
	methodName := getMethodName()
	return &mock.Call{
		Parent:          parent,
		Method:          methodName,
		Arguments:       methodArguments,
		ReturnArguments: make([]interface{}, 0),
		Repeatability:   0,
		WaitFor:         nil,
		RunFn:           nil,
		PanicMsg:        nil,
	}
}

func getMethodName() string {
	// get the calling function's name
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		panic("Couldn't get the caller information")
	}
	functionPath := runtime.FuncForPC(pc).Name()
	// Next four lines are required to use GCCGO function naming conventions.
	// For Ex:  github_com_docker_libkv_store_mock.WatchTree.pN39_github_com_docker_libkv_store_mock.Mock
	// uses interface information unlike golang github.com/docker/libkv/store/mock.(*Mock).WatchTree
	// With GCCGO we need to remove interface information starting from pN<dd>.
	re := regexp.MustCompile("\\.pN\\d+_")
	if re.MatchString(functionPath) {
		functionPath = re.Split(functionPath, -1)[0]
	}
	parts := strings.Split(functionPath, ".")
	return parts[len(parts)-1]
}
