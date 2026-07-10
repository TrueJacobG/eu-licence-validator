package main

import "github.com/TrueJacobG/eu-licence-validator/core/validator"

func IsValid(plate, country string) bool {
	return validator.IsValid(plate, country)
}
