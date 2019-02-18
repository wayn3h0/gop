package zh

import (
	"github.com/wayn3h0/gop/i18n"
	_ "github.com/wayn3h0/gop/i18n/zh-hans"
)

func init() {
	zh, _ := i18n.LookupLanguage("zh")
	zh_hans, _ := i18n.LookupLanguage("zh-hans")

	for _, country := range i18n.AllCountries() {
		if name, ok := country.Name.Get(zh_hans); ok {
			country.Name.Set(zh, name)
		}
		if aliases, ok := country.Aliases.Get(zh_hans); ok {
			country.Aliases.Set(zh, aliases...)
		}
	}

	for _, currency := range i18n.AllCurrencies() {
		if name, ok := currency.Name.Get(zh_hans); ok {
			currency.Name.Set(zh, name)
		}
	}
}
