# finident

[![Go Reference](https://pkg.go.dev/badge/github.com/robfordww/finident.svg)](https://pkg.go.dev/github.com/robfordww/finident)
[![Go Report Card](https://goreportcard.com/badge/github.com/robfordww/finident)](https://goreportcard.com/report/github.com/robfordww/finident)

`finident` is a Go library for finance teams that need accurate Legal Entity
Identifier (LEI) validation, ISO 6166 ISIN checksum recalculation, ISO 10962 CFI
verification, and deterministic ESMA UTI generation. Use it to embed
regulatory-grade identifier checks and trade-reporting helpers into Go
microservices, batch pipelines, or custom tooling. The project is distributed
under the MIT License.

## Features

- Go-native validation for Legal Entity Identifiers (LEI, ISO 17442) including
  checksum verification
- International Securities Identification Number (ISIN, ISO 6166) parsing and
  recalculated check digits
- Classification of Financial Instruments (CFI, ISO 10962) validation and bulk
  combination generation based on ESMA guidance
- ESMA CPMI-IOSCO-aligned Unique Transaction Identifier (UTI) hashing helpers
  with time + randomness fallback for reliable uniqueness
- Utility helpers for mod-97 checksum math and CFI tooling so you can compose
  higher-level finance services quickly

## Installation

```bash
go get github.com/robfordww/finident
```

The module requires Go 1.20 or later. Import it in your project and call the
helpers you need:

```go
import "github.com/robfordww/finident"
```

## Quick start

```go
package main

import (
    "fmt"

    "github.com/robfordww/finident"
)

func main() {
    if ok, err := finident.ValidateLEI("5493004W1IPC50878Z34"); !ok {
        fmt.Println("LEI failed:", err)
    }

    if ok, err := finident.ValidateISIN("US0378331005"); !ok {
        fmt.Println("ISIN failed:", err)
    }

    if finident.IsValidCFI("ESVTOB") {
        fmt.Println("CFI looks good")
    }
}
```

## Legal Entity Identifier Validation in Go

`ValidateLEI` mirrors the ISO 17442 rule set for Legal Entity Identifiers. It
checks string length, enforces the reserved `00` characters, restricts the
character set to uppercase alphanumerics, and verifies the mod-97 checksum. Pair
it with `CalculateChecksum` if you need to derive compliant suffix digits when
issuing new LEIs programmatically.

## Generate ESMA-Compliant UTIs

To produce CPMI-IOSCO/ESMA-compliant Unique Transaction Identifiers, use
`GenerateUTI` when you already possess the 32-character suffix, or
`GenerateUTIFromParts` to hash trade metadata and generate the suffix
deterministically. When no metadata is available, the helper combines UTC
nanoseconds with a cryptographically random 32-bit prefix so concurrently running
systems still produce unique UTIs.

## Classification of Financial Instruments (CFI) Support

`IsValidCFI` and `GenCFICombinations` encode the ESMA mapping for ISO 10962
codes. You can quickly validate single CFIs, enumerate every supported
instrument class for discovery workflows, or cache the combinations to power
form builders and selection widgets.

## Common Use Cases

- Enrich trade reporting pipelines that must publish ESMA-aligned UTIs and LEIs
- Validate inbound securities master data before it hits downstream services
- Generate CFI picklists for compliance portals and audit dashboards
- Prototype Go microservices that expose validator APIs for finance operations

## API overview

- `ValidateLEI(lei string) (bool, error)` checks 20-character LEIs for uppercase
  alphanumerics, reserved `00` positions at indices 4–5, and the ISO 17442 mod-97
  checksum, returning a validity flag plus descriptive error.
- `ValidateISIN(isin string) (bool, error)` validates 12-character ISINs by
  enforcing the leading country code letters and recalculating the ISO 6166
  check digit; failures return an explanatory error.
- `IsValidCFI(cfi string) bool` confirms a six-character CFI matches the ESMA
  taxonomy segments stored in the map and returns `true` only for valid
  combinations.
- `GenCFICombinations() []string` materializes every permitted CFI combination
  from the ESMA mapping, yielding several thousand entries as a convenience for
  enumeration tasks.
- `Validatemod97(s string) bool` normalizes the input to uppercase, converts
  letters to their 10–35 numeric equivalents, and reports whether the numeric
  value modulo 97 equals 1.
- `CalculateChecksum(s string) string` computes and zero-pads the two-digit
  string that must be appended so the full value satisfies the mod-97==1 rule
  used by LEIs.
- `GenerateUTI(lei, value string) (string, error)` concatenates a validated LEI
  with a caller-provided value and ensures the result meets the CPMI-IOSCO/ESMA
  guidance (upper-case alphanumerics with a maximum length of 52).
- `GenerateUTIFromParts(lei string, parts ...string) (string, error)` derives the
  value component deterministically (SHA-256 + base32) from trade metadata so
  that callers can produce ESMA-compliant UTIs without hand-crafting the
  32-character suffix.

### Notes on CFI generation

`GenCFICombinations` expands every valid segment defined by the European
Securities and Markets Authority (ESMA). The resulting slice contains several
thousand entries; allocate or stream the values accordingly if you only need a
subset. The source mapping is derived from
[`ESMA's official taxonomy`](https://www.esma.europa.eu/file/20301/download?token=6K3VKc5m).

## FAQ

**How do I calculate an LEI checksum in Go?** Use `CalculateChecksum` to compute
the two-character suffix that makes a base string pass the ISO 17442 mod-97
check.

**How can I ensure UTI uniqueness across services?** `GenerateUTIFromParts`
hashes ordered metadata with SHA-256 and falls back to a random 32-bit prefix
plus UTC nanoseconds, so services running in parallel will still produce unique
identifiers.

**Does finident help with CFI discovery?** `GenCFICombinations` expands every
valid ISO 10962 segment so you can materialize dropdowns, filter lists, or cache
the data for analytics jobs.

## Contributing & development

- Run tests with `go test ./...` (set `GOCACHE` if your environment requires a
  writable cache).
- Please open an issue or pull request if you discover new instrument classes or
  additional validation rules we should cover.

## License

This project is available under the MIT License. See `LICENSE.md` for details.
