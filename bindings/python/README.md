# eu-licence-validator (Python)

European licence plate validation for Python, powered by a shared Go/WebAssembly core.

## Install

```bash
pip install eu-licence-validator
```

## Usage

```python
from eu_licence_validator import is_valid

is_valid("WPI 1234X", "PL")   # True
is_valid("AA-123-SS", "FR")   # False
```
