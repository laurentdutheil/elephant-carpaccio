package money

import (
	"fmt"
	"math"
	"strconv"
)

type Decimal int64

func (d Decimal) Multiply(other Decimal) Decimal {
	return Decimal(math.Round(float64(d) * float64(other) * math.Pow10(-2)))
}

func (d Decimal) Divide(other Decimal) Decimal {
	return Decimal(math.Round(float64(d) / float64(other) * math.Pow10(2)))
}

func (d Decimal) String() string {
	units := d / 100
	formattedUnits := addThousandsSeparators(units)
	decimals := d % 100
	return fmt.Sprintf("%v.%02d", formattedUnits, decimals)
}

func addThousandsSeparators(units Decimal) string {
	formattedUnits := strconv.FormatInt(int64(units), 10)
	for i := len(formattedUnits); i > 3; {
		i -= 3
		formattedUnits = formattedUnits[:i] + "," + formattedUnits[i:]
	}
	return formattedUnits
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	units := d / 100
	decimals := d % 100
	return []byte(fmt.Sprintf("%d.%02d", units, decimals)), nil
}

func (d Decimal) Floor() Decimal {
	return Decimal(math.Round(float64(d)*math.Pow10(-2)) * math.Pow10(2))
}

func (d Decimal) Ceil() Decimal {
	floor := d.Floor()
	if d == floor {
		return d
	}
	return floor + Decimal(100)
}
