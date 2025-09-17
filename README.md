# finident

`finident` is a Go toolbox for working with common financial identifiers. It
currently focuses on:

- Legal Entity Identifier (LEI) validation (ISO 17442)
- International Securities Identification Number (ISIN) validation (ISO 6166)
- Classification of Financial Instruments (CFI) validation and combination
  generation (ISO 10962)

The library is distributed under the MIT License.

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

## Contributing & development

- Run tests with `go test ./...` (set `GOCACHE` if your environment requires a
  writable cache).
- Please open an issue or pull request if you discover new instrument classes or
  additional validation rules we should cover.

## License

This project is available under the MIT License. See `LICENSE.md` for details.
