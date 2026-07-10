package eulicencevalidator

import "github.com/eu-licence-validator/core/validator"

func IsValid(plate, countryCode string) bool {
	return validator.IsValid(plate, countryCode)
}
