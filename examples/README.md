# Examples

One minimal, runnable example per language. Each example imports the
**published** package (not the local source) to prove the end-to-end packaging
works.

> These will work once the first release is published to each registry. Until
> then they are illustrative of the final API. Final package names settle with
> the Phase 3 language bindings.

The API is identical in every language:

```
isValid(plate, countryCode) -> bool
```

| Language | Run |
| --- | --- |
| Python | `python examples/python/example.py` |
| Node.js | `node examples/node/example.js` |
| Go | `go run examples/go/example.go` |
| Ruby | `ruby examples/ruby/example.rb` |
| Java | see `examples/java/` |

## Install the package first

```bash
# Python
pip install eu-licence-validator

# Node.js
npm install @truejacobg/eu-licence-validator

# Ruby
gem install eu-licence-validator

# Java
# add the com.github.truejacobg:eu-licence-validator dependency to your pom.xml

# Go
go get github.com/TrueJacobG/eu-licence-validator/bindings/go
```
