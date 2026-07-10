# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Go validation core with support for Poland (`PL`), Germany (`DE`), and
  France SIV (`FR`).
- TinyGo WASI WebAssembly exports (`alloc`, `validate`, `dealloc`) in
  `core/main.go`.
- Unified test data in `test_cases.json`, consumed by the Go core test suite.
- `Makefile` with `build-wasm`, `test-all`, `test-core`, and `clean` targets.
- CI workflow (build + test) and guarded dev-publish / release workflows.
- Project documentation: README, CONTRIBUTING, CODE_OF_CONDUCT, SECURITY,
  CHANGELOG, LICENSE (MIT).
- Examples scaffolding for Node, Python, Go, Ruby, and Java.

### Notes
- Country rules are intentionally permissive placeholders; they will be
  tightened per country in subsequent iterations.
- Language wrappers (Phase 3) are not yet implemented; publish steps skip
  until the corresponding binding exists.

[Unreleased]: https://github.com/TrueJacobG/eu-licence-validator/compare/HEAD
