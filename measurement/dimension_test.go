package measurement

import (
	"encoding/json"
	"testing"

	testing2 "github.com/wayn3h0/gop/testing"
)

func TestDimension(t *testing.T) {
	data := map[string]map[DimensionUnit]string{
		"101.5 mm": {
			Millimeter: "101.5mm",
			Centimeter: "10.15cm",
			Meter:      "0.1015m",
		},
		"102.5   cm": {
			Millimeter: "1025mm",
			Centimeter: "102.5cm",
			Meter:      "1.025m",
		},
		"103.5m": {
			Millimeter: "103500mm",
			Centimeter: "10350cm",
			Meter:      "103.5m",
		},
	}
	for k, v := range data {
		d := MustParseDimension(k)
		for k2, v2 := range v {
			d.Convert(k2)
			testing2.AssertEqual(t, d.String(), v2)
		}
	}

	data2 := []string{
		"101.5mm",
		"102.5cm",
		"103.5m",
	}
	for _, v := range data2 {
		d1 := MustParseDimension(v)
		d, err := json.Marshal(d1)
		if err != nil {
			t.Fatal(err)
		}
		var d2 Dimension
		err = json.Unmarshal(d, &d2)
		if err != nil {
			t.Fatal(err)
		}
		testing2.AssertEqual(t, d1.String(), (&d2).String())
	}
}
