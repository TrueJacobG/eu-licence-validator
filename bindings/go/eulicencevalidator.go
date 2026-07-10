package eulicencevalidator

import "github.com/TrueJacobG/eu-licence-validator/core/validator"

func IsValid(plate, countryCode string) bool {
	return validator.IsValid(plate, countryCode)
}
