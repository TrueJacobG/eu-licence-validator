import json
import os
from pathlib import Path

import pytest

from eu_licence_validator import is_valid

_CASES_PATH = Path(__file__).resolve().parents[3] / "test_cases.json"


@pytest.fixture(scope="module")
def cases():
    with open(_CASES_PATH) as f:
        data = json.load(f)
    assert data, "test_cases.json is empty"
    return data


@pytest.mark.parametrize("case", [{}])  # replaced dynamically below
def test_cases_placeholder(case):
    pass


def _ids(cases):
    return [f"{c['country']}:{c['plate']}" for c in cases]


def test_all_cases(cases):
    failures = []
    for c in cases:
        got = is_valid(c["plate"], c["country"])
        if got != c["expected"]:
            failures.append(
                f"is_valid({c['plate']!r}, {c['country']!r}) = {got}, want {c['expected']}"
            )
    assert not failures, "\n".join(failures)
