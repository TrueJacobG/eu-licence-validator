module example-test

go 1.26.5

replace github.com/TrueJacobG/eu-licence-validator/bindings/go => /Users/jakubgradzewicz/DEV/go/eu-licence-validator/bindings/go

replace github.com/TrueJacobG/eu-licence-validator/core => /Users/jakubgradzewicz/DEV/go/eu-licence-validator/core

require github.com/TrueJacobG/eu-licence-validator/bindings/go v0.0.0-00010101000000-000000000000

require github.com/TrueJacobG/eu-licence-validator/core v0.0.0 // indirect
