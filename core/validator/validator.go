package validator

import (
	"regexp"
	"strings"
)

type validatorFunc func(string) bool

var validators = map[string]validatorFunc{
	"PL": validatePL,
	"DE": validateDE,
	"FR": validateFR,
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

var plRe = regexp.MustCompile(`^[A-Z]{2,3}[\s-]?[0-9A-Z]{3,5}$`)

func validatePL(plate string) bool {
	return plRe.MatchString(plate)
}

var deRe = regexp.MustCompile(`^[A-Z]{1,3}[\s-]?[A-Z]{1,2}[\s-]?\d{1,4}[EH]?$`)

func validateDE(plate string) bool {
	return deRe.MatchString(plate)
}

var frRe = regexp.MustCompile(`^[A-Z]{2}[\s-]?\d{3}[\s-]?[A-Z]{2}$`)

func validateFR(plate string) bool {
	if !frRe.MatchString(plate) {
		return false
	}
	clean := stripSeparators(plate)
	if len(clean) != 7 {
		return false
	}
	for _, g := range []string{clean[0:2], clean[5:7]} {
		if strings.ContainsAny(g, "IOU") {
			return false
		}
		if g == "SS" || g == "WW" {
			return false
		}
	}
	return true
}
