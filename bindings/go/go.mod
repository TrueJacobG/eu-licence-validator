module github.com/TrueJacobG/eu-licence-validator/bindings/go

go 1.22

require github.com/TrueJacobG/eu-licence-validator/core v0.0.1

// For local development inside the monorepo: use the local core module copy.
// This replace directive is ignored by the public Go proxy but ensures `go test`
// works when running inside the monorepo (developer workflow / CI here).
replace github.com/TrueJacobG/eu-licence-validator/core => ../../core
