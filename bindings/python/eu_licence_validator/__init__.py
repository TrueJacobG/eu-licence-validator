"""European licence plate validation (Python binding).

Loads the shared Go-compiled WebAssembly core once and exposes
``is_valid(plate, country_code) -> bool``.
"""

from __future__ import annotations

import os
from functools import lru_cache
from typing import Tuple

from wasmtime import Engine, ExitTrap, Instance, Linker, Module, Store, WasiConfig

__all__ = ["is_valid", "isValid"]

_WASM_PATH = os.path.join(os.path.dirname(__file__), "wasm", "core.wasm")


@lru_cache(maxsize=1)
def _instance() -> Tuple[Store, Instance]:
    engine = Engine()
    store = Store(engine)
    store.set_wasi(WasiConfig())
    module = Module.from_file(engine, _WASM_PATH)
    linker = Linker(engine)
    linker.define_wasi()
    instance = linker.instantiate(store, module)
    exports = instance.exports(store)

    start = exports["_start"]
    if start is not None:
        try:
            start(store)
        except (ExitTrap, SystemExit, RuntimeError):
            pass
    return store, instance


def _write(store: Store, instance: Instance, data: bytes) -> Tuple[int, int]:
    exports = instance.exports(store)
    alloc = exports["alloc"]
    ptr = alloc(store, len(data))
    if ptr == 0 and len(data) > 0:
        raise MemoryError("wasm alloc returned 0")
    mem = exports["memory"]
    raw = mem.data_ptr(store)
    for i, b in enumerate(data):
        raw[ptr + i] = b
    return ptr, len(data)


def is_valid(plate: str, country_code: str) -> bool:
    """Return True if ``plate`` is a valid licence plate for ``country_code``."""
    store, instance = _instance()
    exports = instance.exports(store)

    plate_bytes = plate.encode("utf-8")
    country_bytes = country_code.encode("utf-8")

    p_ptr, p_len = _write(store, instance, plate_bytes)
    c_ptr, c_len = _write(store, instance, country_bytes)

    validate = exports["validate"]
    result = validate(store, p_ptr, p_len, c_ptr, c_len)

    dealloc = exports["dealloc"]
    dealloc(store, p_ptr)
    dealloc(store, c_ptr)
    return result == 1


isValid = is_valid
