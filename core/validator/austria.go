package validator

import "regexp"

var atRe = regexp.MustCompile(`^([A-Z]{1,2})\s?(\d{3,5})\s?([A-Z]?)$`)

func validateAT(plate string) bool {
	m := atRe.FindStringSubmatch(plate)
	if m == nil {
		return false
	}
	return onlyAlpha(m[1])
}
