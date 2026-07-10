SHELL := /bin/bash

TINYGO := tinygo
GO     := go
NODE   := npm
PYTHON := python3
RUBY   ?= $(shell for r in ruby /opt/homebrew/opt/ruby/bin/ruby $$HOME/.rbenv/shims/ruby; do command -v $$r >/dev/null 2>&1 && $$r -e 'require "wasmtime"' >/dev/null 2>&1 && { echo $$r; break; }; done)
MVN    := mvn

CORE_DIR  := core
WASM_BIN  := $(CORE_DIR)/core.wasm

WASM_DESTS := \
	bindings/node/wasm \
	bindings/python/eu_licence_validator/wasm \
	bindings/ruby/wasm \
	bindings/java/src/main/resources

.PHONY: all build-wasm distribute-wasm test-all test-core test-node test-python test-ruby test-java test-go clean

all: build-wasm

build-wasm: distribute-wasm

distribute-wasm:
	@if [ ! -f $(WASM_BIN) ]; then \
		echo ">> $(WASM_BIN) not found, building with TinyGo..."; \
		cd $(CORE_DIR) && $(TINYGO) build -target wasi -no-debug -o core.wasm .; \
	fi
	@echo ">> Distributing core.wasm to language bindings..."
	@for dir in $(WASM_DESTS); do \
		mkdir -p $$dir; \
		cp $(WASM_BIN) $$dir/core.wasm; \
		echo "   copied -> $$dir/core.wasm"; \
	done

test-all: test-core test-node test-python test-ruby test-java test-go
	@echo ">> All language test suites complete."

test-core:
	@echo ">> Testing Go core..."
	cd $(CORE_DIR) && $(GO) test ./...

test-node:
	@echo ">> Testing Node.js binding..."
	@if [ -f bindings/node/package.json ]; then \
		cd bindings/node && $(NODE) ci && $(NODE) test; \
	else echo "   (skipped: bindings/node not set up)"; fi

test-python:
	@echo ">> Testing Python binding..."
	@if [ -f bindings/python/pyproject.toml ] || [ -f bindings/python/setup.py ]; then \
		cd bindings/python && $(PYTHON) -m pip install -q -e ".[dev]" && $(PYTHON) -m pytest -q; \
	else echo "   (skipped: bindings/python not set up)"; fi

test-ruby:
	@echo ">> Testing Ruby binding..."
	@if ls bindings/ruby/*.gemspec >/dev/null 2>&1; then \
		cd bindings/ruby && $(RUBY) -Ilib -e 'Dir["test/**/*_test.rb"].each { |f| require File.expand_path(f) }'; \
	else echo "   (skipped: bindings/ruby not set up)"; fi

test-java:
	@echo ">> Testing Java binding..."
	@if [ -f bindings/java/pom.xml ]; then \
		cd bindings/java && $(MVN) -q test; \
	else echo "   (skipped: bindings/java not set up)"; fi

test-go:
	@echo ">> Testing native Go binding..."
	@if [ -f bindings/go/go.mod ]; then \
		cd bindings/go && $(GO) test ./...; \
	else echo "   (skipped: bindings/go not set up)"; fi

clean:
	rm -f $(WASM_BIN)
	@for dir in $(WASM_DESTS); do rm -f $$dir/core.wasm; done
	@echo ">> Cleaned wasm artifacts."
