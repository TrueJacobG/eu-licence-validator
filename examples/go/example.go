package main

import (
	"fmt"

	valid "github.com/TrueJacobG/eu-licence-validator/bindings/go"
)

func main() {
	plates := []struct {
		plate   string
		country string
	}{
		{"WPI 1234X", "PL"},
		{"B-AB 1234", "DE"},
		{"AA-123-AB", "FR"},
		{"AA-123-SS", "FR"},
		{"WPI 1234X", "XX"},
	}

	for _, p := range plates {
		fmt.Printf("isValid(%q, %q) = %v\n", p.plate, p.country, valid.IsValid(p.plate, p.country))
	}
}
