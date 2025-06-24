# Financial identifiers validation library
Functions for validating Legal Entity Identifier (LEI), as described in ISO 17442,  International Securities Identification Numbers (ISIN) codes, and CFI codes (ISO 10962) generator and validator.

MIT License Copyright (c) 2019, robfordww


# FUNCTIONS
```
func CalculateChecksum(s string) string
    CalculateChecksum takes a string and returns the next two characters
    that, when appended to the string, results in a "stringvalue mod 97 ==
    1"

func GenCFICombinations() []string
    GenCFICombinations returns a list of all valid CFIs. This validator is
    based on ESMAs CFI list published here
    [https://www.esma.europa.eu/file/20301/download?token=6K3VKc5m](https://www.esma.europa.eu/document/firds-cfi-validations)

func IsValidCFI(cfi string) bool
    IsValidCFI returns true if the CFI string is a valid CFI code, and false
    otherwise. This validator is based on ESMAs CFI list published here
    [https://www.esma.europa.eu/file/20301/download?token=6K3VKc5m](https://www.esma.europa.eu/document/firds-cfi-validations)

func ValidateISIN(isin string) (bool, error)
    ValidateISIN takes an ISIN (ISO 6166) as string and validates the
    checkdigit

func ValidateLEI(lei string) (bool, error)
    ValidateLEI takes a possible LEI string as input and returns a boolean
    value error message. For valid LEIs, the bool == true and error == nil,
    otherwise bool == false and error != nil

func Validatemod97(s string) bool
    Validatemod97 takes a string as parameter and returns true if mod 97 of
    the string, interpreted as a number, returns 1. Letters A-Z are
    converted to numbers 10-34


```
