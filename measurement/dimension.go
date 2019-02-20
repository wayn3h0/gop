package measurement

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/wayn3h0/gop/decimal"
	"github.com/wayn3h0/gop/errors"
)

// DimensionUnit represents the unit of dimension.
type DimensionUnit int

// Dimension Units.
const (
	Millimeter DimensionUnit = 1 // base
	Centimeter DimensionUnit = 10
	Meter      DimensionUnit = 1000
)

var (
	// default dimension unit
	DefaultDimensionUnit = Centimeter
)

// Dimension represents a dimension information.
// The value of dimension will reserves to MinimumDimensionUnit.
type Dimension struct {
	unit  DimensionUnit
	value *decimal.Decimal
}

func (d *Dimension) ensureInitialized() {
	if d.value == nil {
		d.unit = DefaultDimensionUnit
		d.value = new(decimal.Decimal)
	}
}

// Unit returns the unit of dimension.
func (d *Dimension) Unit() DimensionUnit {
	return d.unit
}

// Value returns the float64 value of dimension and a indicator whether the value is exact.
func (d *Dimension) Value() (float64, bool) {
	return d.value.Float64()
}

// IsZero reports whether the value is zero.
func (d *Dimension) IsZero() bool {
	d.ensureInitialized()
	return d.value.IsZero()
}

// String returns the formatted string.
func (d *Dimension) String() string {
	d.ensureInitialized()
	switch d.unit {
	case Millimeter:
		return d.value.String() + "mm"
	case Centimeter:
		return d.value.String() + "cm"
	case Meter:
		return d.value.String() + "m"
	}

	panic(errors.Newf("measurement: unknown dimension unit `%d`", d.unit))
}

// Copy sets x to y and return x.
func (x *Dimension) Copy(y *Dimension) *Dimension {
	x.ensureInitialized()
	y.ensureInitialized()
	x.unit = y.unit
	x.value.Copy(y.value)
	return x
}

// Convert converts the dimension with given unit.
func (d *Dimension) Convert(unit DimensionUnit) *Dimension {
	d.ensureInitialized()
	if d.unit != unit {
		if !d.value.IsZero() {
			src := new(decimal.Decimal).SetInt64(int64(d.unit))
			dst := new(decimal.Decimal).SetInt64(int64(unit))
			d.value.Mul(src).Div(dst)
		}
		d.unit = unit
	}
	return d
}

// RoundUp rounds the value up to integer.
func (d *Dimension) RoundUp() *Dimension {
	d.value.RoundUp(0)
	return d
}

// Cmp compares x and y and returns:
// -1 if x < y
//  0 if x == y
// +1 if x > y
func (x *Dimension) Cmp(y *Dimension) int {
	x.ensureInitialized()
	y.ensureInitialized()
	yval := new(Dimension).Copy(y).Convert(x.unit).value
	return x.value.Cmp(yval)
}

// Add set x to x+y and return x.
func (x *Dimension) Add(y *Dimension) *Dimension {
	x.ensureInitialized()
	y.ensureInitialized()
	yval := new(Dimension).Copy(y).Convert(x.unit).value
	x.value.Add(yval)
	return x
}

// Sub sets x to x-y and return x.
func (x *Dimension) Sub(y *Dimension) *Dimension {
	x.ensureInitialized()
	y.ensureInitialized()
	yval := new(Dimension).Copy(y).Convert(x.unit).value
	x.value.Sub(yval)
	return x
}

// Mul sets x to x*y and return x.
func (x *Dimension) Mul(y float64) *Dimension {
	x.ensureInitialized()
	x.value.Mul(decimal.New(y))
	return x
}

// Div sets x to x/y and return x.
func (x *Dimension) Div(y float64) *Dimension {
	x.ensureInitialized()
	x.value.Div(decimal.New(y))
	return x
}

// NewDimension returns a new weight.
func NewDimension(value float64, unit DimensionUnit) (*Dimension, error) {
	val := decimal.New(value)
	if val.Sign() < 0 {
		return nil, errors.Newf("measurement: dimension value `%f' is invalid", value)
	}
	return &Dimension{
		unit:  unit,
		value: val,
	}, nil
}

// MustNewDimension is similar to NewDimension but panics if has error.
func MustNewDimension(value float64, unit DimensionUnit) *Dimension {
	d, err := NewDimension(value, unit)
	if err != nil {
		panic(err)
	}
	return d
}

var (
	_DimensionPattern = regexp.MustCompile(`^(\d+\.?\d*)(\s*)(mm|cm|m)?$`)
)

// ParseDimension returns a new weight by parsing string.
func ParseDimension(str string) (*Dimension, error) {
	s := strings.ToLower(strings.Replace(str, " ", "", -1))
	matches := _DimensionPattern.FindStringSubmatch(s)
	if len(matches) != 4 {
		return nil, errors.Newf("measurement: dimension string %q is invalid", str)
	}
	value, _ := strconv.ParseFloat(matches[1], 64)
	var unit DimensionUnit
	switch matches[3] {
	case "mm":
		unit = Millimeter
	case "cm":
		unit = Centimeter
	case "m":
		unit = Meter
	default:
		unit = DefaultDimensionUnit
	}

	return NewDimension(value, unit)
}

// MustParseDimension is similar to ParseDimension but panics if has error.
func MustParseDimension(str string) *Dimension {
	d, err := ParseDimension(str)
	if err != nil {
		panic(err)
	}
	return d
}

// MarshalJSON marshals to JSON data.
// Implements Marshaler interface.
func (d *Dimension) MarshalJSON() ([]byte, error) {
	d.ensureInitialized()
	return []byte(strconv.Quote(d.String())), nil
}

// UnmarshalJSON unmarshas from JSON data.
// Implements Unmarshaler interface.
func (d *Dimension) UnmarshalJSON(data []byte) error {
	str, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	d2, err := ParseDimension(str)
	if err != nil {
		return err
	}
	d.Copy(d2)
	return nil
}

// Dimensions represents a sortable collection of dimension.
type Dimensions []*Dimension

// implements sort.Interface.
func (ds Dimensions) Len() int {
	return len(ds)
}

// implements sort.Interface.
func (ds Dimensions) Less(i, j int) bool {
	return ds[i].Cmp(ds[j]) < 0
}

// implements sort.Interface.
func (ds Dimensions) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

func (ds Dimensions) Sort() {
	sort.Sort(ds)
}
