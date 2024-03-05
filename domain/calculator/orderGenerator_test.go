package calculator_test

import (
	. "elephant_carpaccio/domain/money"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"regexp"
	"runtime"
	"strings"
	"testing"

	. "elephant_carpaccio/domain/calculator"
)

func TestGenerateOrder(t *testing.T) {
	t.Run("should generate an order with fixed Discount and fixed State", func(t *testing.T) {
		orderGenerator := NewOrderGenerator(NewOrderRandomizer())

		order := orderGenerator.GenerateOrder(No.Discount(), AL.State())
		receipt := order.Compute()

		assert.Equal(t, AL.State(), order.State)
		assert.Equal(t, No.Discount(), receipt.Discount)
	})

	t.Run("should generate a ramdom nbItems between Decimal(1) included and Decimal(10000) excluded", func(t *testing.T) {
		spyOrderRandom := NewSpyOrderRandom(NewOrderRandomizer())

		orderGenerator := NewOrderGenerator(spyOrderRandom)
		order := orderGenerator.GenerateOrder(No.Discount(), UT.State())

		assert.GreaterOrEqual(t, order.NumberOfItems, Decimal(1))
		assert.Less(t, order.NumberOfItems, Decimal(10000))
		spyOrderRandom.AssertCalled(t, "RandDecimal", Decimal(1), Decimal(10000))
	})

	t.Run("should generate a random discount order between minimal Discount amount and maximal Discount amount", func(t *testing.T) {
		spyOrderRandom := NewSpyOrderRandom(NewOrderRandomizer())
		orderGenerator := NewOrderGenerator(spyOrderRandom)
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
		orderGenerator := NewOrderGenerator(spyOrderRandom)

		order := orderGenerator.GenerateOrder(No.Discount(), nil)

		assert.NotNil(t, order.State)
		spyOrderRandom.AssertCalled(t, "RandState")
	})

	t.Run("should pick a discount level at random when argument is nil", func(t *testing.T) {
		spyOrderRandom := NewSpyOrderRandom(NewOrderRandomizer())
		orderGenerator := NewOrderGenerator(spyOrderRandom)

		order := orderGenerator.GenerateOrder(nil, AL.State())
		receipt := order.Compute()

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
