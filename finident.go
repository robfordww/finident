// Package finident is a toolbox for various Financial identifiers
// MIT License
// Copyright (c) 2019 robfordww
package finident

import (
	"fmt"
	"strconv"
)

const (
	charshift = 55
	numshift  = 48
)

// LeiError returned on error
type LeiError error

// ValidateLEI takes a possible LEI string as input and returns a boolean value error message.
// For valid LEIs, the bool == true and error == nil, otherwise bool == false and error != nil
func ValidateLEI(lei string) (bool, error) {
	// Validate length
	if len(lei) != 20 {
		return false, LeiError(fmt.Errorf("Wrong length of LEI code, %v bytes", len(lei)))
	}
	// Validate reserved characters
	if lei[4] != '0' || lei[5] != '0' {
		return false, LeiError(fmt.Errorf("Reserved charaters 5 & 6 are not zero"))
	}
	// Charaterset A-Z
	for _, r := range lei {
		if !((r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false, LeiError(fmt.Errorf("Invalid character %q in LEI-code", r))
		}
	}
	// Consider only upper case chars
	r := Validatemod97(lei)
	if !r {
		return r, LeiError(fmt.Errorf("Checksum failed"))
	}
	return r, nil
}

// ValidateISIN takes an ISIN (ISO 6166) as string and validates the checkdigit
func ValidateISIN(isin string) (bool, error) {
	if l := len(isin); l != 12 {
		return false, fmt.Errorf("Invalid length of ISIN (length:%v)", l)
	}
	if !(isA2Z(isin[0]) && isA2Z(isin[1])) {
		return false, fmt.Errorf("Two first characters must be letters A-Z (%v)", isin[0:2])
	}
	var isinchecksum int
	var poslogic = 1
	for i := range isin {
		v := isin[len(isin)-i-1] // reverse scan string
		if isA2Z(v) {
			fd := int(v-charshift) / 10 // first digit
			sd := int(v-charshift) % 10 // second digit
			if poslogic == 0 {
				isinchecksum += fd + sumOfDigits(2*sd)
			} else {
				isinchecksum += sumOfDigits(2*fd) + sd
			}
			poslogic ^= 1 // flip bit to account for a character representing 2 digits
		} else if v >= '0' && v <= '9' {
			if poslogic == 0 {
				isinchecksum += sumOfDigits(int(v-numshift) * 2)
			} else {
				isinchecksum += sumOfDigits(int(v - numshift))
			}
		} else {
			return false, fmt.Errorf("Invalid character in ISIN string: %v", v)
		}
		poslogic ^= 1
	}
	if isinchecksum%10 != 0 {
		return false, fmt.Errorf("Checksumdigit failed: %v", isinchecksum)
	}
	return true, nil
}

// Validatemod97 takes a string as parameter and returns true if mod 97 of the
// string, interpreted as a number, returns 1. Letters A-Z are converted to
// numbers 10-34
func Validatemod97(s string) bool {
	if mod97([]byte(s)) != 1 {
		return false
	}
	return true
}

// CalculateChecksum takes a string and returns the next two characters that,
// when appended to the string, results in a "stringvalue mod 97 == 1"
func CalculateChecksum(s string) string {
	return strconv.Itoa(98 - ((100 * int(mod97([]byte(s)))) % 97))
}

// checks if byte is a-Z
func isA2Z(b byte) bool {
	if b >= 'A' && b <= 'z' {
		return true
	}
	return false
}

// Sum of digits. Also works for negative numbers; -123 => -6
func sumOfDigits(i int) (sum int) {
	for ; i != 0; i /= 10 {
		sum += int(i % 10)
	}
	return sum
}

func mod97(s []byte) int64 {
	var checksum int64
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			checksum *= 100
			checksum += int64(r) - charshift
		} else if r >= 'a' && r <= 'z' {
			checksum *= 100
			checksum += int64(r - charshift)
		} else if r >= '0' && r <= '9' {
			checksum *= 10
			checksum += int64(r) - numshift
		} else {
			panic("undefined char for mod 97")
		}
		if (i+1)%8 == 0 {
			checksum = checksum % 97
		}
	}
	checksum = checksum % 97
	return checksum
}
