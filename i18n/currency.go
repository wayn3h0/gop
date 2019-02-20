package i18n

import (
	"sort"
	"strings"
)

// Currency represents a currency information. ISO 4217
type Currency struct {
	AlphaCode   string // ISO alpha currency code
	NumericCode string // ISO numeric currency code
	Name        String
}

// Eaual reports whether two currencies are same.
// It compares the alpha code.
func (x *Currency) Equal(y *Currency) bool {
	return strings.EqualFold(x.AlphaCode, y.AlphaCode)
}

// Currencies represents a sortable collection of currency.
type Currencies []*Currency

func (cs Currencies) Len() int {
	return len(cs)
}

func (cs Currencies) Less(i, j int) bool {
	return cs[i].AlphaCode < cs[j].AlphaCode
}

func (cs Currencies) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs Currencies) Sort() {
	sort.Sort(cs)
}

var (
	mapOfAlphaCodeToCurrency   map[string]*Currency
	mapOfNumericCodeToCurrency map[string]*Currency
	listOfAllCurrencies        Currencies
)

func init() {
	mapOfAlphaCodeToCurrency = make(map[string]*Currency)
	mapOfNumericCodeToCurrency = make(map[string]*Currency)
	listOfAllCurrencies = make(Currencies, len(listOfCurrencyCodes))
	for i, codes := range listOfCurrencyCodes {
		alphaCode := codes[0]
		numericCode := codes[1]
		country := &Currency{
			AlphaCode:   alphaCode,
			NumericCode: numericCode,
			Name:        make(String),
		}

		mapOfAlphaCodeToCurrency[alphaCode] = country
		mapOfNumericCodeToCurrency[numericCode] = country
		listOfAllCurrencies[i] = country
	}

	listOfAllCurrencies.Sort()
}

// AllCurrencies returns the list of all countries.
func AllCurrencies() Currencies {
	return listOfAllCurrencies
}

// LookupCurrency returns the country by given code.
func LookupCurrency(code string) (*Currency, bool) {
	c := strings.ToUpper(code)
	// alpha code & numeric code
	if len(c) == 3 {
		if v, ok := mapOfAlphaCodeToCurrency[c]; ok {
			return v, true
		}
		if v, ok := mapOfNumericCodeToCurrency[c]; ok {
			return v, true
		}
	}

	return nil, false
}
