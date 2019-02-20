package i18n

import (
	"sort"
	"strings"
)

// Country represents a country information. ISO 3166-1
type Country struct {
	Alpha2Code  string // ISO alpha-2 country code
	Alpha3Code  string // ISO alpha-3 country code
	NumericCode string // ISO numeric country code
	Name        String
	Aliases     StringArray
}

// Eaual reports whether two countries are same.
// It compares the alpha-2 code.
func (x *Country) Equal(y *Country) bool {
	return strings.EqualFold(x.Alpha2Code, y.Alpha2Code)
}

// Countries represents a collection of Country.
type Countries []*Country

func (cs Countries) Len() int {
	return len(cs)
}

func (cs Countries) Less(i, j int) bool {
	return cs[i].Alpha2Code < cs[j].Alpha2Code
}

func (cs Countries) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs Countries) Sort() {
	sort.Sort(cs)
}

var (
	mapOfAlpha2CodeToCountry  map[string]*Country
	mapOfAlpha3CodeToCountry  map[string]*Country
	mapOfNumericCodeToCountry map[string]*Country
	listOfAllCountries        Countries
)

func init() {
	mapOfAlpha2CodeToCountry = make(map[string]*Country)
	mapOfAlpha3CodeToCountry = make(map[string]*Country)
	mapOfNumericCodeToCountry = make(map[string]*Country)
	listOfAllCountries = make(Countries, len(listOfCountryCodes))
	for i, codes := range listOfCountryCodes {
		alpha2Code := codes[0]
		alpha3Code := codes[1]
		numericCode := codes[2]
		country := &Country{
			Alpha2Code:  alpha2Code,
			Alpha3Code:  alpha3Code,
			NumericCode: numericCode,
			Name:        make(String),
			Aliases:     make(StringArray),
		}

		mapOfAlpha2CodeToCountry[alpha2Code] = country
		mapOfAlpha3CodeToCountry[alpha3Code] = country
		mapOfNumericCodeToCountry[numericCode] = country
		listOfAllCountries[i] = country
	}

	listOfAllCountries.Sort()
}

// AllCountries returns the list of all countries.
func AllCountries() Countries {
	return listOfAllCountries
}

// LookupCountry returns the country by given code.
func LookupCountry(code string) (*Country, bool) {
	c := strings.ToUpper(code)
	// alpha-2 code
	if len(c) == 2 {
		if v, ok := mapOfAlpha2CodeToCountry[c]; ok {
			return v, true
		}
	}
	// alpha-3 & numeric code
	if len(c) == 3 {
		if v, ok := mapOfAlpha3CodeToCountry[c]; ok {
			return v, true
		}
		if v, ok := mapOfNumericCodeToCountry[c]; ok {
			return v, true
		}
	}

	return nil, false
}

// SearchCountries returns the countries by given keyword with specified language.
func SearchCountries(lang *Language, keyword string) Countries {
	kw := strings.ToLower(keyword)
	found := make(map[string]*Country)
	for _, country := range listOfAllCountries {
		// compare name
		if name, ok := country.Name.Get(lang); ok {
			if strings.Contains(strings.ToLower(name), kw) {
				found[country.Alpha2Code] = country
			}
		}

		// compare aliases
		if aliases, ok := country.Aliases.Get(lang); ok {
			str := strings.Join(aliases, StringArrayValueSeparator)
			if strings.Contains(strings.ToLower(str), kw) {
				found[country.Alpha2Code] = country
			}
		}
	}

	countries := make(Countries, len(found))
	index := 0
	for _, country := range found {
		countries[index] = country
		index++
	}

	return countries
}
