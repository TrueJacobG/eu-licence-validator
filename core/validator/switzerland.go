package validator

import "regexp"

var chRe = regexp.MustCompile(`^([A-Z]{2})\s?(\d{1,6})$`)

func validateCH(plate string) bool {
	return chRe.MatchString(plate)
}
