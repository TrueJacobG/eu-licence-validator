package validator

import (
	"strings"
)

type validatorFunc func(string) bool

var validators = map[string]validatorFunc{
	"PL": validatePL,
	"DE": validateDE,
	"FR": validateFR,
	"IT": validateIT,
	"ES": validateES,
	"NL": validateNL,
	"BE": validateBE,
	"PT": validatePT,
	"AT": validateAT,
	"CH": validateCH,
}

func IsValid(plate, country string) bool {
	fn, ok := validators[strings.ToUpper(strings.TrimSpace(country))]
	if !ok {
		return false
	}
	return fn(strings.ToUpper(strings.TrimSpace(plate)))
}

func stripSeparators(s string) string {
	return strings.NewReplacer(" ", "", "\t", "", "-", "").Replace(s)
}

func onlyDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func onlyAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < 'A' || c > 'Z' {
			return false
		}
	}
	return true
}
