package zh_hant

import (
	"github.com/wayn3h0/gop/i18n"
)

func init() {
	lang, _ := i18n.LookupLanguage("zh-hant")

	for code, name := range mapOfCountryAlpha2CodeToName {
		if country, ok := i18n.LookupCountry(code); ok {
			country.Name.Set(lang, name)
		}
	}
	for code, aliases := range mapOfCountryAlpha2CodeToAliases {
		if country, ok := i18n.LookupCountry(code); ok {
			country.Aliases.Set(lang, aliases...)
		}
	}

	for code, name := range mapOfCurrencyAlphaCodeToName {
		if currency, ok := i18n.LookupCurrency(code); ok {
			currency.Name.Set(lang, name)
		}
	}
}
