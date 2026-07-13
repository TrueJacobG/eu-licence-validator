package validator

import "regexp"

var itRe = regexp.MustCompile(`^([A-Z]{2})\s?(\d{3})[,\s]?([A-Z]{2})$`)

func validateIT(plate string) bool {
	return itRe.MatchString(plate)
}
