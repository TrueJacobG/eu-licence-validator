package validator

import "regexp"

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
