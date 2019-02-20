package measurement

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/wayn3h0/gop/decimal"
	"github.com/wayn3h0/gop/errors"
)

// VolumeUnit represents the unit of volume.
type VolumeUnit int

// Volume Units.
const (
	CubicMillimeter VolumeUnit = 1 // base
	CubicCentimeter VolumeUnit = 1000
	CubicMeter      VolumeUnit = 1000000000
)

var (
	// default volume unit
	DefaultVolumeUnit = CubicMeter
)

// Volume represents a volume information.
type Volume struct {
	unit  VolumeUnit
	value *decimal.Decimal
}

func (v *Volume) ensureInitialized() {
	if v.value == nil {
		v.unit = DefaultVolumeUnit
		v.value = new(decimal.Decimal)
	}
}

// Unit returns the unit of volume.
func (v *Volume) Unit() VolumeUnit {
	return v.unit
}

// Value returns the float64 value of volume and a indicator whether the value is exact.
func (v *Volume) Value() (float64, bool) {
	return v.value.Float64()
}

// IsZero reports whether the value is zero.
func (v *Volume) IsZero() bool {
	v.ensureInitialized()
	return v.value.IsZero()
}

// String returns the formatted string.
func (v *Volume) String() string {
	v.ensureInitialized()
	switch v.unit {
	case CubicMillimeter:
		return v.value.String() + "mm3"
	case CubicCentimeter:
		return v.value.String() + "cm3"
	case CubicMeter:
		return v.value.String() + "m3"
	}

	panic(errors.Newf("measurement: unknown volume unit `%d`", v.unit))
}

// Copy set x to y and return x.
func (x *Volume) Copy(y *Volume) *Volume {
	x.ensureInitialized()
	y.ensureInitialized()
	x.unit = y.unit
	x.value.Copy(y.value)
	return x
}

// Convert converts the volume with given unit.
func (v *Volume) Convert(unit VolumeUnit) *Volume {
	v.ensureInitialized()
	if v.unit != unit {
		if !v.value.IsZero() {
			src := new(decimal.Decimal).SetInt64(int64(v.unit))
			dst := new(decimal.Decimal).SetInt64(int64(unit))
			v.value.Mul(src).Div(dst)
		}
		v.unit = unit

	}
	return v
}

// RoundUp rounds the value up to integer.
func (v *Volume) RoundUp() *Volume {
	v.value.RoundUp(0)
	return v
}

// Cmp compares x and y and returns:
// -1 if x < y
//  0 if x == y
// +1 if x > y
func (x *Volume) Cmp(y *Volume) int {
	x.ensureInitialized()
	y.ensureInitialized()
	yval := new(Volume).Copy(y).Convert(x.unit).value
	return x.value.Cmp(yval)
}

// Add sets x to x+y and returns x.
func (x *Volume) Add(y *Volume) *Volume {
	x.ensureInitialized()
	y.ensureInitialized()
	yval := new(Volume).Copy(y).Convert(x.unit).value
	x.value.Add(yval)
	return x
}

// Sub sets x to x-y and returns x.
func (x *Volume) Sub(y *Volume) *Volume {
	x.ensureInitialized()
	y.ensureInitialized()
	yval := new(Volume).Copy(y).Convert(x.unit).value
	x.value.Sub(yval)
	return x
}

// Mul sets x to x*y and returns x.
func (x *Volume) Mul(y float64) *Volume {
	x.ensureInitialized()
	x.value.Mul(decimal.New(y))
	return x
}

// Div sets x to quotient of x/y and returns x.
func (x *Volume) Div(y float64) *Volume {
	x.ensureInitialized()
	x.value.Div(decimal.New(y))
	return x
}

// Weight returns the dimensional (volumetric) weight in kg.
// Description: https://en.wikipedia.org/wiki/Dimensional_weight
func (v *Volume) Weight(metricFactor int) *Weight {
	v.ensureInitialized()
	val := new(Volume).Copy(v).Convert(CubicCentimeter).value
	w := &Weight{
		unit:  Kilogram,
		value: val,
	}
	return w.Div(float64(metricFactor))
}

// NewVolume returns a new volume.
func NewVolume(value float64, unit VolumeUnit) (*Volume, error) {
	val := decimal.New(value)
	if val.Sign() < 0 {
		return nil, errors.Newf("measurement: volume value `%f' is invalid", value)
	}
	return &Volume{
		unit:  unit,
		value: val,
	}, nil
}

// MustNewVolume is similar to NewVolume but panics if has error.
func MustNewVolume(value float64, unit VolumeUnit) *Volume {
	v, err := NewVolume(value, unit)
	if err != nil {
		panic(err)
	}
	return v
}

// NewVolumeFromDimensions returns a new volumes by dimensions.
func NewVolumeFromDimensions(length, width, height *Dimension) (*Volume, error) {
	if length == nil || length.IsZero() {
		return nil, errors.New("measurement: length is invalid")
	}
	if width == nil || width.IsZero() {
		return nil, errors.New("measurement: width is invalid")
	}
	if height == nil || height.IsZero() {
		return nil, errors.New("measurement: height is invalid")
	}
	l := new(Dimension).Copy(length).Convert(Millimeter).value
	w := new(Dimension).Copy(width).Convert(Millimeter).value
	h := new(Dimension).Copy(height).Convert(Millimeter).value
	volume := &Volume{
		unit:  CubicMillimeter,
		value: l.Mul(w).Mul(h),
	}
	return volume.Convert(DefaultVolumeUnit), nil
}

// MustNewVolumeFromDimensions is similar to NewVolumeFromDimension but panics if has error.
func MustNewVolumeFromDimensions(length, width, height *Dimension) *Volume {
	v, err := NewVolumeFromDimensions(length, width, height)
	if err != nil {
		panic(err)
	}
	return v
}

var (
	_VolumePattern = regexp.MustCompile(`^(\d+\.?\d*)(\s*)(mm3|cm3|m3)?$`)
)

// ParseVolume returns a new volume by parsing string.
// Default: m3
func ParseVolume(str string) (*Volume, error) {
	s := strings.ToLower(strings.Replace(str, " ", "", -1))
	matches := _VolumePattern.FindStringSubmatch(s)
	if len(matches) != 4 {
		return nil, errors.Newf("measurement: volume string %q is invalid", str)
	}
	value, _ := strconv.ParseFloat(matches[1], 64)
	var unit VolumeUnit
	switch matches[3] {
	case "mm3":
		unit = CubicMillimeter
	case "cm3":
		unit = CubicCentimeter
	case "m3":
		unit = CubicMeter
	default:
		unit = DefaultVolumeUnit
	}

	return NewVolume(value, unit)
}

// MustParseVolume is similar to ParseVolume but panics if has error.
func MustParseVolume(str string) *Volume {
	v, err := ParseVolume(str)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalJSON marshals to JSON data.
// Implements Marshaler interface.
func (v *Volume) MarshalJSON() ([]byte, error) {
	v.ensureInitialized()
	return []byte(strconv.Quote(v.String())), nil
}

// UnmarshalJSON unmarshas from JSON data.
// Implements Unmarshaler interface.
func (v *Volume) UnmarshalJSON(data []byte) error {
	str, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	v2, err := ParseVolume(str)
	if err != nil {
		return err
	}
	v.Copy(v2)
	return nil
}
