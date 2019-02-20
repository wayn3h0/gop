package i18n

import (
	"strings"
)

// String represents a multilingual string.
type String map[string]string

// SupportedLanguages returns supported languages.
func (s String) SupportedLanguages() Languages {
	if len(s) == 0 {
		return nil
	}

	langs := make(Languages, len(s))
	index := 0
	for code, _ := range s {
		if lang, ok := LookupLanguage(code); ok {
			langs[index] = lang
			index += 1
		}
	}

	return langs
}

// IsEmpty reports whether the string is empty.
func (s String) IsEmpty() bool {
	return len(s) == 0
}

// Get returns the value with given language.
func (s String) Get(lang *Language) (string, bool) {
	if v, ok := s[lang.Code]; ok {
		return v, true
	}

	return "", false
}

// Set sets the value with given language.
func (s String) Set(lang *Language, value string) {
	s[lang.Code] = value
}

var (
	StringArrayValueSeparator = "|"
)

// StringArray represents an array of multilingual string.
type StringArray String

// Get returns the value of multilingual string with given language.
func (sa StringArray) Get(lang *Language) ([]string, bool) {
	if vals, ok := sa[lang.Code]; ok {
		return strings.Split(vals, StringArrayValueSeparator), true
	}

	return nil, false
}

// Set sets the value of multilingual string with given language.
func (sa StringArray) Set(lang *Language, values ...string) {
	vals := strings.Join(values, StringArrayValueSeparator)
	sa[lang.Code] = vals
}
