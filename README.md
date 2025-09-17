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

## Available helpers

```go
func ValidateLEI(lei string) (bool, error)
func ValidateISIN(isin string) (bool, error)
func IsValidCFI(cfi string) bool
func GenCFICombinations() []string
func Validatemod97(s string) bool
func CalculateChecksum(s string) string
```

### Function reference

- `ValidateLEI` verifies that the input has 20 characters, uses the allowed
  alphanumeric set, keeps reserved positions five and six as `00`, and passes
  the ISO 17442 mod-97 checksum. The boolean return indicates validity and the
  error provides context when validation fails.
- `ValidateISIN` ensures the string is exactly 12 characters, starts with two
  uppercase letters, and satisfies the ISO 6166 checksum calculation. It returns
  a boolean flag plus an explanatory error when the check digit or input
  structure is invalid.
- `IsValidCFI` checks whether a six-character CFI string matches one of the
  ESMA-defined attribute combinations held in the embedded lookup table.
- `GenCFICombinations` expands the ESMA taxonomy into every valid six-character
  CFI code and returns the full set as a slice. Expect several thousand entries
  depending on the current mapping.
- `Validatemod97` converts alphanumeric characters to their numeric
  representations (letters map to 10–35) and returns `true` when the value mod
  97 equals 1. Lowercase letters are handled by normalizing to uppercase before
  evaluation.
- `CalculateChecksum` outputs the two-character string that must be appended to
  the provided base so that the combined value satisfies the mod-97 checksum
  rule (result padded to two digits as required by ISO 17442).

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
