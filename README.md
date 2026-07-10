# eu-licence-validator

[![CI](https://github.com/TrueJacobG/eu-licence-validator/actions/workflows/ci.yml/badge.svg)](https://github.com/TrueJacobG/eu-licence-validator/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A single source of truth for **European licence-plate validation**, delivered as
native-feeling packages for **Node.js / TypeScript**, **Python**, **Ruby**,
**Java**, and **Go**.

The validation logic is written **once** in Go and compiled to a tiny
WebAssembly (WASI) binary with [TinyGo]. The binary is bundled inside each
language package, so end users just call a normal function — they never know they
are running Wasm. The native Go package imports the core logic directly (no
Wasm) for maximum performance.

| Language                | Package                                                  | Distribution          |
| ----------------------- | -------------------------------------------------------- | --------------------- |
| TypeScript / JavaScript | `@truejacobg/eu-licence-validator`                       | npm / GitHub Packages |
| Python                  | `eu-licence-validator`                                   | PyPI                  |
| Ruby                    | `eu-licence-validator`                                   | RubyGems              |
| Java                    | `com.github.truejacobg:eu-licence-validator`             | Maven Central         |
| Go                      | `github.com/TrueJacobG/eu-licence-validator/bindings/go` | Go module             |

> The bindings are landing in Phase 3. Until the first public release, the
> packages above may not yet exist on the public registries.

## Supported countries

| Code | Country      | Status |
| ---- | ------------ | ------ |
| `PL` | Poland       | ✅     |
| `DE` | Germany      | ✅     |
| `FR` | France (SIV) | ✅     |
| `IT` | Italy        | ✅     |
| `ES` | Spain        | ✅     |
| `NL` | Netherlands  | ✅     |
| `BE` | Belgium      | ✅     |
| `PT` | Portugal     | ✅     |
| `AT` | Austria      | ✅     |
| `CH` | Switzerland  | ✅     |

More countries are tracked in [CONTRIBUTING.md](#) — contributions welcome.

## Install

> Replace the version placeholder with the published version once the first
> release is out.

**TypeScript / JavaScript** (Node, Bun, Deno, Vite, browsers)

```bash
npm install @truejacobg/eu-licence-validator
```

**Python**

```bash
pip install eu-licence-validator
```

**Ruby**

```bash
gem install eu-licence-validator
```

**Java (Maven)**

```xml
<dependency>
  <groupId>io.github.truejacobg</groupId>
  <artifactId>eu-licence-validator</artifactId>
  <version>VERSION</version>
</dependency>
```

**Go**

```bash
go get github.com/TrueJacobG/eu-licence-validator/bindings/go
```

## Usage

Every binding exposes the same API:

```ts
// TypeScript / JavaScript (Node, Bun, Deno, Vite, browsers)
import { init, isValid } from "@truejacobg/eu-licence-validator";

await init(); // initialize the Wasm runtime once
isValid("WPI 1234X", "PL"); // true  (sync after init)
isValid("AA-123-SS", "FR"); // false
isValid("B-AB 1234", "DE"); // true

// Types are included:
// function isValid(plate: string, countryCode: string): boolean
// function init(): Promise<void>
```

```python
from eu_licence_validator import is_valid

is_valid("WPI 1234X", "PL")   # True
is_valid("AA-123-SS", "FR")   # False
is_valid("B-AB 1234", "DE")  # True
```

See [`examples/`](examples) for a runnable example in each language, including
Deno and Vite/browser setups.

## How it works

```
                ┌─────────────────────────────────────────────┐
                │              core/ (Go logic)                │
                │   validator.go  ·  regex per EU country      │
                └──────────────────────┬──────────────────────┘
                       (TinyGo, WASI)  │  (direct import)
        ┌──────────────┬──────────────┴───────────────┬──────────────┐
        ▼              ▼                               ▼              ▼
   js (wasm)      python (wasm)     ruby (wasm)   java (wasm)    go (native)
   bundled        bundled           bundled        bundled        imports core
   core.wasm     core.wasm         core.wasm      core.wasm      (no Wasm)
```

- `core/validator.go` — the real validation logic (pure Go, no I/O).
- `core/main.go` — TinyGo WASI exports: `alloc`, `validate`, `dealloc`.
- Each binding instantiates `core.wasm`, calls `_start` once, then uses the
  `alloc` → write bytes → `validate` → `dealloc` protocol.
- The Go binding skips Wasm entirely and calls `IsValid(...)` directly.

## Development

### Prerequisites

| Tool                         | Version |
| ---------------------------- | ------- |
| Go                           | 1.22+   |
| [TinyGo](https://tinygo.org) | latest  |
| Node.js                      | 20+     |
| Python                       | 3.10+   |
| Ruby                         | 3.0+    |
| Java (Maven)                 | 17+     |

### Build & test

```bash
make build-wasm   # compile core -> core.wasm, copy into every binding
make test-all     # run tests across every language binding
make test-core    # Go core tests only
make clean        # remove generated wasm artifacts
```

The Go core can be developed and tested without TinyGo:

```bash
cd core && go test ./...
```

### Repository layout

```
eu-licence-validator/
├── Makefile
├── test_cases.json          # unified test data for every language
├── core/                     # Go + WebAssembly core
├── bindings/                # js · python · ruby · java · go
├── examples/                 # one runnable example per language
└── .github/workflows/        # ci · dev-publish · release
```

## Python - Pypi

[pypi](https://pypi.org/project/eu-licence-validator/)

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request.
By participating you agree to abide by the [Code of Conduct](CODE_OF_CONDUCT.md).

## License

[MIT](LICENSE) © TrueJacobG
