package measurement

import (
	"encoding/json"
	"testing"

	testing2 "github.com/wayn3h0/gop/testing"
)

func TestVolume(t *testing.T) {
	data := map[string]map[VolumeUnit]string{
		"101.5 mm3": {
			CubicMillimeter: "101.5mm3",
			CubicCentimeter: "0.1015cm3",
			CubicMeter:      "0.0000001015m3",
		},
		"102.5cm3": {
			CubicMillimeter: "102500mm3",
			CubicCentimeter: "102.5cm3",
			CubicMeter:      "0.0001025m3",
		},
		"103.5  m3": {
			CubicMillimeter: "103500000000mm3",
			CubicCentimeter: "103500000cm3",
			CubicMeter:      "103.5m3",
		},
	}
	for k, v := range data {
		d := MustParseVolume(k)
		for k2, v2 := range v {
			d.Convert(k2)
			testing2.AssertEqual(t, d.String(), v2)
		}
	}

	data2 := []string{
		"101.5mm3",
		"102.5cm3",
		"103.5m3",
	}
	for _, v := range data2 {
		v1 := MustParseVolume(v)
		d, err := json.Marshal(v1)
		if err != nil {
			t.Fatal(err)
		}
		var v2 Volume
		err = json.Unmarshal(d, &v2)
		if err != nil {
			t.Fatal(err)
		}
		testing2.AssertEqual(t, v1.String(), (&v2).String())
	}

	// test *Volume.Weight func
	data3 := map[string]map[int]string{
		"102.5cm3": {
			5000: "0.021kg",
			6000: "0.018kg",
			8000: "0.013kg",
		},
	}
	for k, v := range data3 {
		volume := MustParseVolume(k)
		for k2, v2 := range v {
			weight := volume.Weight(k2)
			weight.Convert(Gram).RoundUp().Convert(Kilogram)
			testing2.AssertEqual(t, weight.String(), v2)
		}
	}
}
