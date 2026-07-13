package validator

import "regexp"

var esRe = regexp.MustCompile(`^(\d{4})\s?([A-Z]{3})$|^([A-Z]{1,2})\s?(\d{4})\s?([A-Z]{2})$`)

func validateES(plate string) bool {
	return esRe.MatchString(plate)
}
