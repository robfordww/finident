// Financial identifiers validation library
// MIT License
// Copyright (c) 2016 robfordww
package finident

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// LeiError returned on error
type LeiError error

// ValidateLEI takes a possible LEI string as input and returns a boolean value error message.
// For valid LEIs, the bool == true and error == nil, otherwise bool == false and error != nil
func ValidateLEI(lei string) (bool, error) {
	// Validate length
	if len(lei) != 20 {
		return false, LeiError(errors.New("Wrong length of LEI code"))
	}
	// Validate reserved characters
	if lei[4] != '0' || lei[5] != '0' {
		return false, LeiError(errors.New("Reserved charaters 5 & 6 are not zero"))
	}
	// Charaterset A-Z
	for _, r := range lei {
		if !((r >= 'A' && r <= 'z') || (r >= '0' && r <= '9')) {
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

// Validatemod97 takes a string as parameter and returns true if mod 97 of the
// string, interpreted as a number, returns 1. Letters A-Z are converted to
// numbers 10-34
func Validatemod97(s string) bool {
	if mod97(s) != 1 {
		return false
	}
	return true
}

// CalculateChecksum takes a string and returns the next two characters that
// ,when appended to the string, results in a "stringvalue mod 97 == 1"
func CalculateChecksum(s string) string {
	return strconv.Itoa(98 - ((100 * int(mod97(s))) % 97))
}

func mod97(s string) int64 {
	upperlei := strings.ToUpper(s)
	const charshift = 55
	const numshift = 48
	var checksum int64
	for i, r := range upperlei {
		if r >= 'A' && r <= 'Z' {
			checksum *= 100
			checksum += int64(r) - charshift
		} else if r >= 'a' && r <= 'z' {
			checksum += int64(unicode.ToUpper(r)) - charshift
			panic("wtf")
		} else if r >= '0' && r <= '9' {
			checksum *= 10
			checksum += int64(r) - numshift
		} else {
			panic("wtf")
		}
		//fmt.Println("D:", checksum)
		if (i+1)%8 == 0 {
			checksum = checksum % 97
		}
	}
	checksum = checksum % 97
	return checksum
}
