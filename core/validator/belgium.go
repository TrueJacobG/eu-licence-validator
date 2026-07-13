package validator

import "regexp"

var beRe = regexp.MustCompile(`^([A-Z]{3})[-\s]?(\d{3})$|^(\d)-([A-Z]{3})-(\d{3})$|^(\d{4})[-\s]?([A-Z]{3})$`)

func validateBE(plate string) bool {
	return beRe.MatchString(plate)
}
