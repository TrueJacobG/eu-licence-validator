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

var plVoivodeships = map[byte]bool{
	'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true,
	'H': true, 'J': true, 'K': true, 'L': true, 'N': true, 'O': true,
	'P': true, 'R': true, 'S': true, 'T': true, 'W': true, 'Z': true,
}

var plRe = regexp.MustCompile(`^([A-Z]{2,3})\s?([0-9A-Z]{3,5})$`)

func validatePL(plate string) bool {
	m := plRe.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	region := m[1]
	seq := m[2]

	if len(region) < 2 || !plVoivodeships[region[0]] {
		return false
	}

	if len(seq) < 4 || len(seq) > 5 {
		return false
	}

	digitCount := 0
	for _, c := range seq {
		if c >= '0' && c <= '9' {
			digitCount++
		}
	}
	if digitCount < 1 {
		return false
	}

	return true
}

var deRe = regexp.MustCompile(`^([A-Z]{1,3})[\s-]?([A-Z]{1,2})[\s-]?(\d{1,4})([EH]?)$`)

func validateDE(plate string) bool {
	m := deRe.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	district := m[1]
	middle := m[2]

	clean := stripSeparators(plate)
	if len(clean) > 8 {
		return false
	}

	if len(district) == 0 || len(middle) == 0 {
		return false
	}

	return true
}

var frRe = regexp.MustCompile(`^([A-Z]{2})[\s-]?(\d{3})[\s-]?([A-Z]{2})$`)

func validateFR(plate string) bool {
	m := frRe.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	first := m[1]
	last := m[3]

	for _, g := range []string{first, last} {
		if strings.ContainsAny(g, "IOU") {
			return false
		}
		if g == "SS" || g == "WW" {
			return false
		}
	}
	return true
}

var itRe = regexp.MustCompile(`^([A-Z]{2})\s?(\d{3})[,\s]?([A-Z]{2})$`)

func validateIT(plate string) bool {
	return itRe.MatchString(plate)
}

var esRe = regexp.MustCompile(`^(\d{4})\s?([A-Z]{3})$|^([A-Z]{1,2})\s?(\d{4})\s?([A-Z]{2})$`)

func validateES(plate string) bool {
	return esRe.MatchString(plate)
}

var nlRe = regexp.MustCompile(`^([A-Z]{2})[-\s]?(\d{2})[-\s]?(\d{2})$|^(\d{2})[-\s]?([A-Z]{2})[-\s]?(\d{2})$|^(\d{1,2})[-\s]?([A-Z]{3})[-\s]?(\d)$`)

func validateNL(plate string) bool {
	return nlRe.MatchString(plate)
}

var beRe = regexp.MustCompile(`^([A-Z]{3})[-\s]?(\d{3})$|^(\d)-([A-Z]{3})-(\d{3})$|^(\d{4})[-\s]?([A-Z]{3})$`)

func validateBE(plate string) bool {
	return beRe.MatchString(plate)
}

var ptRe = regexp.MustCompile(`^([A-Z]{2})[-\s]?(\d{2})[-\s]?([A-Z]{2})$`)

func validatePT(plate string) bool {
	return ptRe.MatchString(plate)
}

var atRe = regexp.MustCompile(`^([A-Z]{1,2})\s?(\d{3,5})\s?([A-Z]?)$`)

func validateAT(plate string) bool {
	m := atRe.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	return onlyAlpha(m[1])
}

var chRe = regexp.MustCompile(`^([A-Z]{2})\s?(\d{1,6})$`)

func validateCH(plate string) bool {
	return chRe.MatchString(plate)
}
