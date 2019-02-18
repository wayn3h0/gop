package measurement

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/wayn3h0/gop/decimal"
	"github.com/wayn3h0/gop/errors"
)

// WeightUnit represents the unit of weight.
type WeightUnit int

// Weight Units.
const (
	Gram     WeightUnit = 1 // base
	Kilogram WeightUnit = 1000
	Tonne    WeightUnit = 1000000
)

var (
	// default weight unit
	DefaultWeightUnit = Kilogram
)

// Weight represents a weight information.
// The value of weight will reserves to MinimumWeightUnit.
type Weight struct {
	unit  WeightUnit
	value *decimal.Decimal
}

func (w *Weight) ensureInitialized() {
	if w.value == nil {
		w.unit = DefaultWeightUnit
		w.value = new(decimal.Decimal)
	}
}

// Unit returns the unit of weight.
func (w *Weight) Unit() WeightUnit {
	return w.unit
}

// Value returns the float64 value of weight and a indicator whether the value is exact.
func (w *Weight) Value() (float64, bool) {
	return w.value.Float64()
}

// IsZero reports whether the value is zero.
func (w *Weight) IsZero() bool {
	w.ensureInitialized()
	return w.value.IsZero()
}

// String returns the formatted string.
func (w *Weight) String() string {
	w.ensureInitialized()
	switch w.unit {
	case Gram:
		return w.value.String() + "g"
	case Kilogram:
		return w.value.String() + "kg"
	case Tonne:
		return w.value.String() + "t"
	}

	panic(errors.Newf("measurement: unknown weight unit `%d`", w.unit))
}

// Copy set x to y and return x.
func (x *Weight) Copy(y *Weight) *Weight {
	x.ensureInitialized()
	y.ensureInitialized()
	x.unit = y.unit
	x.value.Copy(y.value)
	return x
}

// Convert converts the weight to given unit.
func (w *Weight) Convert(unit WeightUnit) *Weight {
	w.ensureInitialized()
	if w.unit != unit {
		if !w.value.IsZero() {
			src := new(decimal.Decimal).SetInt64(int64(w.unit))
			dst := new(decimal.Decimal).SetInt64(int64(unit))
			w.value.Mul(src).Div(dst)
		}
		w.unit = unit
	}
	return w
}

// RoundUp rounds the value up to integer.
func (w *Weight) RoundUp() *Weight {
	w.value.RoundUp(0)
	return w
}

// Cmp compares x and y and returns:
// -1 if x < y
//  0 if x == y
// +1 if x > y
func (x *Weight) Cmp(y *Weight) int {
	x.ensureInitialized()
	y.ensureInitialized()
	yval := new(Weight).Copy(y).Convert(x.unit).value
	return x.value.Cmp(yval)
}

// Align aligns to base weight.
// Example: 1.25kg align by 0.5kg, result: 1.5kg.
func (w *Weight) Align(base *Weight) *Weight {
	stepping := new(Weight).Copy(base).Convert(w.unit).value
	if w.value.Cmp(stepping) <= 0 {
		w.value.Copy(stepping)
		return w
	}
	w.value.Quo(stepping).RoundUp(0).Mul(stepping)
	return w
}

// Add set x to x+y and return x.
func (x *Weight) Add(y *Weight) *Weight {
	x.ensureInitialized()
	y.ensureInitialized()
	yval := new(Weight).Copy(y).Convert(x.unit).value
	x.value.Add(yval)
	return x
}

// Sub sets x to x-y and return x.
func (x *Weight) Sub(y *Weight) *Weight {
	x.ensureInitialized()
	y.ensureInitialized()
	yval := new(Weight).Copy(y).Convert(x.unit).value
	x.value.Sub(yval)
	return x
}

// Mul sets x to x*y and return x.
func (x *Weight) Mul(y float64) *Weight {
	x.ensureInitialized()
	x.value.Mul(decimal.New(y))
	return x
}

// Div sets x to x/y and return x.
func (x *Weight) Div(y float64) *Weight {
	x.ensureInitialized()
	x.value.Div(decimal.New(y))
	return x
}

// NewWeight returns a new weight.
func NewWeight(value float64, unit WeightUnit) (*Weight, error) {
	val := decimal.New(value)
	if val.Sign() < 0 {
		return nil, errors.Newf("measurement: weight value `%f' is invalid", value)
	}
	return &Weight{
		unit:  unit,
		value: val,
	}, nil
}

// MustNewWeight is similar to NewWeight but panics if has error.
func MustNewWeight(value float64, unit WeightUnit) *Weight {
	w, err := NewWeight(value, unit)
	if err != nil {
		panic(err)
	}
	return w
}

var (
	_WeightPattern = regexp.MustCompile(`^(\d+\.?\d*)(\s*)(g|kg|t)?$`)
)

// ParseWeight returns a new weight by parsing string.
func ParseWeight(str string) (*Weight, error) {
	s := strings.ToLower(strings.Replace(str, " ", "", -1))
	matches := _WeightPattern.FindStringSubmatch(s)
	if len(matches) != 4 {
		return nil, errors.Newf("measurement: weight string %q is invalid", str)
	}
	value, _ := strconv.ParseFloat(matches[1], 64)
	var unit WeightUnit
	switch matches[3] {
	case "g":
		unit = Gram
	case "kg":
		unit = Kilogram
	case "t":
		unit = Tonne
	default:
		unit = DefaultWeightUnit
	}

	return NewWeight(value, unit)
}

// MustParseWeight is similar to ParseWeight but panics if has error.
func MustParseWeight(str string) *Weight {
	w, err := ParseWeight(str)
	if err != nil {
		panic(err)
	}
	return w
}

// MarshalJSON marshals to JSON data.
// Implements Marshaler interface.
func (w *Weight) MarshalJSON() ([]byte, error) {
	w.ensureInitialized()
	return []byte(strconv.Quote(w.String())), nil
}

// UnmarshalJSON unmarshas from JSON data.
// Implements Unmarshaler interface.
func (w *Weight) UnmarshalJSON(data []byte) error {
	str, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	w2, err := ParseWeight(str)
	if err != nil {
		return err
	}
	w.Copy(w2)
	return nil
}
