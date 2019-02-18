package i18n

import (
	"sort"
	"strings"
)

// Language represents a language. ISO 639-1.
type Language struct {
	Code       string // ISO code of the language
	NativeName string // native name of language
}

// Equal reports whether two languages are same.
// It compares the code.
func (x *Language) Equal(y *Language) bool {
	return strings.EqualFold(x.Code, y.Code)
}

// Languages represents a sortable collection of Language.
type Languages []*Language

// implements sort.Interface.
func (ls Languages) Len() int {
	return len(ls)
}

// implements sort.Interface.
func (ls Languages) Less(i, j int) bool {
	return ls[i].Code < ls[j].Code
}

// implements sort.Interface.
func (ls Languages) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}

// implements sort.Interface.
func (ls Languages) Sort() {
	sort.Sort(ls)
}

// Reverse returns the reverse order for languages.
func (ls Languages) Reverse() {
	sort.Sort(sort.Reverse(ls))
}

var (
	mapOfLanguageCodeToLanguage map[string]*Language
	listOfAllLanguages          Languages
)

func init() {
	mapOfLanguageCodeToLanguage = make(map[string]*Language)
	listOfAllLanguages = make(Languages, len(mapOfLanguageCodeToNativeName))
	index := 0
	for code, name := range mapOfLanguageCodeToNativeName {
		lang := &Language{
			Code:       code,
			NativeName: name,
		}
		mapOfLanguageCodeToLanguage[code] = lang
		listOfAllLanguages[index] = lang
		index++
	}

	listOfAllLanguages.Sort()
}

// AllLanguages returns the list of all languages.
func AllLanguages() Languages {
	return listOfAllLanguages
}

// LookupLanguage returns the language by given code.
func LookupLanguage(code string) (*Language, bool) {
	if lang, ok := mapOfLanguageCodeToLanguage[strings.ToLower(code)]; ok {
		return lang, true
	}
	return nil, false
}
