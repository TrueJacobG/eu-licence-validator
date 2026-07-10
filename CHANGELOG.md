# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Go validation core with support for Poland (`PL`), Germany (`DE`), France SIV
  (`FR`), Italy (`IT`), Spain (`ES`), Netherlands (`NL`), Belgium (`BE`),
  Portugal (`PT`), Austria (`AT`), and Switzerland (`CH`).
- TinyGo WASI WebAssembly exports (`alloc`, `validate`, `dealloc`) in
  `core/main.go` with lazy-initialized memory registry (works without `_start`).
- Unified test data in `test_cases.json` (68 cases), consumed by every language
  binding's test suite.
- `Makefile` with `build-wasm`, `distribute-wasm`, `test-all`, `test-core`,
  `test-js`, `test-python`, `test-ruby`, `test-java`, `test-go`, and `clean`
  targets.
- Language bindings for TypeScript/JavaScript (Node, Bun, Deno, Vite, browsers),
  Python, Ruby, Java, and native Go — all tested against `test_cases.json`.
- Runtime-agnostic JS binding: no `node:wasi` dependency, works in Node, Bun,
  Deno, and browsers/Vite via `WebAssembly.instantiate` with inline WASI shims.
- TypeScript type definitions shipped with the JS package
  (`isValid(plate: string, countryCode: string): boolean`).
- Unified CI/CD pipeline (`.github/workflows/ci.yml`): build-and-test →
  dev-publish (on main) or release-publish (on tags), with 30-minute timeout
  and queued (non-cancelling) concurrency.
- Dev packages published to GitHub Packages (npm, Ruby, Maven) and TestPyPI
  (Python) on every push to main, versioned `0.0.{run_number}`.
- Release packages published to npm, PyPI, RubyGems, and Maven Central on
  `v*` tags.
- Project documentation: README, CONTRIBUTING, CODE_OF_CONDUCT, SECURITY,
  CHANGELOG, LICENSE (MIT).
- Examples for JavaScript (Node/Bun, Deno, Vite), Python, Go, Ruby, and Java.
- Community files: CODEOWNERS, Dependabot (monthly), PR template, issue
  templates (bug report, feature request).
- `.editorconfig` and `.gitattributes` (LF normalization, binary wasm handling).

### Changed
- Tightened PL rules: validates voivodeship first letter, enforces 4-5 char
  sequence with at least 1 digit.
- Tightened DE rules: enforces max 8 chars (without separators), requires
  district + middle letters + 1-4 digits + optional E/H suffix.
- Tightened FR rules: excludes I/O/U from both letter groups, blocks SS and WW.
- Renamed `bindings/node` → `bindings/js` and `examples/node` → `examples/js`
  to reflect the runtime-agnostic JS binding.
- JS API changed from async `isValid()` to `init()` (async, once) + `isValid()`
  (sync) for better DX.
- Consolidated 3 separate workflow files (`ci.yml`, `dev-publish.yml`,
  `release.yml`) into a single `ci.yml` with chained jobs.

[Unreleased]: https://github.com/TrueJacobG/eu-licence-validator/compare/HEAD
