package validator

import (
	"regexp"
	"strings"
)

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
