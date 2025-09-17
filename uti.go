package finident

import (
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	utiMaxLength      = 52
	leiLength         = 20
	utiValueMaxLength = utiMaxLength - leiLength
)

var base32NoPadding = base32.StdEncoding.WithPadding(base32.NoPadding)

// GenerateUTI concatenates the LEI of the generating entity with a caller-supplied
// value and returns a UTI that complies with ESMA's implementation of the CPMI-IOSCO
// UTI Technical Guidance (February 2017). The LEI is validated using ValidateLEI;
// the unique value must contain only A-Z and 0-9 characters and, when appended to
// the LEI, must not exceed 52 characters in total. The returned UTI is always
// upper-case.
func GenerateUTI(lei, value string) (string, error) {
	lei = strings.ToUpper(strings.TrimSpace(lei))
	if ok, err := ValidateLEI(lei); !ok {
		if err == nil {
			err = fmt.Errorf("invalid LEI")
		}
		return "", fmt.Errorf("invalid LEI for UTI generation: %w", err)
	}
	value = strings.ToUpper(strings.TrimSpace(value))
	if len(value) == 0 {
		return "", fmt.Errorf("UTI value must not be empty")
	}
	if len(value) > utiValueMaxLength {
		return "", fmt.Errorf("UTI value exceeds maximum length of %d characters", utiValueMaxLength)
	}
	for _, r := range value {
		if (r < '0' || r > '9') && (r < 'A' || r > 'Z') {
			return "", fmt.Errorf("invalid character %q in UTI value; allowed characters are A-Z and 0-9", r)
		}
	}
	return lei + value, nil
}

// GenerateUTIFromParts derives an ESMA-compliant UTI using the provided LEI and a
// set of inputs that are hashed into the value component. When no additional parts
// are supplied the current UTC time (nanoseconds) is used to guarantee uniqueness.
// The resulting UTI is deterministic for a given LEI and ordered set of parts.
func GenerateUTIFromParts(lei string, parts ...string) (string, error) {
	if len(parts) == 0 {
		parts = []string{strconv.FormatInt(time.Now().UTC().UnixNano(), 10)}
	}
	value := deriveUTIValue(parts)
	return GenerateUTI(lei, value)
}

func deriveUTIValue(parts []string) string {
	h := sha256.New()
	for _, p := range parts {
		h.Write([]byte(p))
		h.Write([]byte{0})
	}
	encoded := base32NoPadding.EncodeToString(h.Sum(nil))
	if len(encoded) > utiValueMaxLength {
		return encoded[:utiValueMaxLength]
	}
	return encoded
}
