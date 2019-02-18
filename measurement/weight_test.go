package measurement

import (
	"encoding/json"
	"testing"

	testing2 "github.com/wayn3h0/gop/testing"
)

func TestWeight(t *testing.T) {
	data := map[string]map[WeightUnit]string{
		"101.5g": {
			Gram:     "101.5g",
			Kilogram: "0.1015kg",
			Tonne:    "0.0001015t",
		},
		"102.5kg": {
			Gram:     "102500g",
			Kilogram: "102.5kg",
			Tonne:    "0.1025t",
		},
		"103.5t": {
			Gram:     "103500000g",
			Kilogram: "103500kg",
			Tonne:    "103.5t",
		},
	}
	for k, v := range data {
		for k2, v2 := range v {
			w := MustParseWeight(k)
			w.Convert(k2)
			testing2.AssertEqual(t, w.String(), v2)
		}
	}

	data2 := []string{
		"101.5g",
		"102.5kg",
		"103.5t",
	}
	for _, v := range data2 {
		w1 := MustParseWeight(v)
		d, err := json.Marshal(w1)
		if err != nil {
			t.Fatal(err)
		}
		var w2 Weight
		err = json.Unmarshal(d, &w2)
		if err != nil {
			t.Fatal(err)
		}
		testing2.AssertEqual(t, w1.String(), (&w2).String())
	}
}
