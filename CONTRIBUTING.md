# Contributing to eu-licence-validator

Thanks for your interest in improving eu-licence-validator! This document
explains how the project is structured and how to add support for a new country
or a new language binding.

## Architecture in one paragraph

The validation logic lives once in Go (`core/validator.go`) and is compiled to a
WASI WebAssembly binary with TinyGo (`make build-wasm`). That binary is bundled
inside every language package in `bindings/`, where a thin host wrapper
instantiates it and exposes a clean `isValid(plate, countryCode)` function.
The native Go binding is the exception: it imports the core logic directly (no
Wasm) for speed. All languages are tested against the same
[`test_cases.json`](../test_cases.json).

## Setup

```bash
git clone https://github.com/TrueJacobG/eu-licence-validator.git
cd eu-licence-validator
make build-wasm
make test-all
```

Required tools: Go 1.22+, [TinyGo](https://tinygo.org) (latest), Node.js 20+,
Python 3.10+, Ruby 3.0+, Java/Maven 17+. See [README](../README.md) for details.

The Go core does not require TinyGo to develop or test:

```bash
cd core && go test ./...
```

## Adding a new country

1. Add a validator function and a compiled `*regexp.Regexp` in
   [`core/validator.go`](../core/validator.go).
2. Register it in the `validators` map under its ISO country code.
3. Add both **valid** and **invalid** entries to
   [`test_cases.json`](../test_cases.json).
4. Run `cd core && go test ./...` then `make build-wasm && make test-all`.
5. Update the supported-countries table in [README](../README.md).

Keep rules as strict as is reasonable and document any edge cases (e.g.
historical formats) as a comment in the regex definition.

## Adding a new language binding

1. Create `bindings/<lang>/` with the standard project manifest
   (`package.json`, `pyproject.toml`, `*.gemspec`, `pom.xml`, `go.mod`, …).
2. Load the bundled `core.wasm` (copied by `make build-wasm`).
3. Implement the host protocol:
   - instantiate the module with a WASI context,
   - call `_start()` once,
   - for each call: `alloc(len)` → write UTF-8 bytes → `validate(ptr, len, ptr2, len2)` (1/0) → `dealloc(ptr)`,
   - re-fetch the memory view after every `alloc` (memory may grow).
4. Expose `isValid(plate, countryCode) -> bool`.
5. Write a test that reads `../../test_cases.json` and loops the assertions
   (no hardcoded cases).
6. Add an example under [`examples/<lang>/`](../examples) and a row in the
   README package table.
7. Wire the publish step into `.github/workflows/dev-publish.yml` and
   `.github/workflows/release.yml`.

## Test data

`test_cases.json` is the **single source of truth**. Never hardcode test cases
in a binding — always read this file. Each entry is:

```json
{ "plate": "WPI 1234X", "country": "PL", "expected": true }
```

## Conventions

- Keep the Go core free of I/O and platform dependencies.
- Do not check in generated `*.wasm` binaries for the bindings (they are built
  by `make build-wasm`); the repo only ships the *source* that produces them.
- Follow the existing code style. Go code passes `go vet` and `gofmt`.
- Commit messages: concise, imperative mood (`add FR SIV rules`).

## Pull requests

- Keep PRs focused — one country, one binding, or one fix at a time.
- Ensure `make build-wasm && make test-all` is green locally.
- Reference any relevant issue (`Closes #123`).
- Use the [pull request template](../.github/pull_request_template.md).

By participating you agree to abide by the [Code of Conduct](CODE_OF_CONDUCT.md).
