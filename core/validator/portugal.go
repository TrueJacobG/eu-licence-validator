package validator

import "regexp"

var ptRe = regexp.MustCompile(`^([A-Z]{2})[-\s]?(\d{2})[-\s]?([A-Z]{2})$`)

func validatePT(plate string) bool {
	return ptRe.MatchString(plate)
}
