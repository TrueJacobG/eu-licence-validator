# Security Policy

## Reporting a Vulnerability

The eu-licence-validator project takes security seriously. If you believe you
have found a security vulnerability, **please do not open a public GitHub
issue**.

Instead, report it privately:

1. Open a **GitHub Security Advisory** via
   the **Report a vulnerability** button on the
   [Security tab](https://github.com/TrueJacobG/eu-licence-validator/security),
   or
2. Email the maintainer at **jakubgradzewicz1309@gmail.com**.

Please include:

- a description of the issue and its potential impact,
- steps or a minimal reproducer,
- the affected version (commit SHA or release tag),
- any suggested mitigation.

You should receive an acknowledgement within **5 business days**. Once the
issue is confirmed, we will coordinate a fix and disclosure timeline with you.

## Scope

Vulnerabilities in the validation logic (e.g. a plate being wrongly accepted
as valid) are **functional bugs**, not security issues — please file those as
normal GitHub issues. Security reports are for issues that could allow code
execution, data leakage, or similar across the host language runtimes (for
example, flaws in the Wasm memory-management protocol).

## Supported Versions

Only the latest released version receives security fixes.

| Version | Supported |
| --- | --- |
| latest `v*` tag | ✅ |
| `main` branch | ✅ (development) |
| older tags | ❌ |
