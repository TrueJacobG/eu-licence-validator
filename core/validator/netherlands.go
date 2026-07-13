package validator

import "regexp"

var nlRe = regexp.MustCompile(`^([A-Z]{2})[-\s]?(\d{2})[-\s]?(\d{2})$|^(\d{2})[-\s]?([A-Z]{2})[-\s]?(\d{2})$|^(\d{1,2})[-\s]?([A-Z]{3})[-\s]?(\d)$`)

func validateNL(plate string) bool {
	return nlRe.MatchString(plate)
}
