package finident

import (
	"regexp"
	"strings"
	"testing"
)

func TestGenerateUTI(t *testing.T) {
	uti, err := GenerateUTI("5493004W1IPC50878Z34", "ABC123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "5493004W1IPC50878Z34ABC123"
	if uti != expected {
		t.Fatalf("expected %s, got %s", expected, uti)
	}

	if _, err := GenerateUTI("INVALID", "ABC123"); err == nil {
		t.Fatalf("expected error for invalid LEI")
	}

	if _, err := GenerateUTI("5493004W1IPC50878Z34", "ABC/"); err == nil {
		t.Fatalf("expected error for invalid characters in value")
	}

	tooLong := strings.Repeat("A", utiValueMaxLength+1)
	if _, err := GenerateUTI("5493004W1IPC50878Z34", tooLong); err == nil {
		t.Fatalf("expected error for value exceeding %d characters", utiValueMaxLength)
	}
}

func TestGenerateUTIFromParts(t *testing.T) {
	parts := []string{"trade-123", "deskA"}
	uti1, err := GenerateUTIFromParts("5493004W1IPC50878Z34", parts...)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	uti2, err := GenerateUTIFromParts("5493004W1IPC50878Z34", parts...)
	if err != nil {
		t.Fatalf("unexpected error on second call: %v", err)
	}
	if uti1 != uti2 {
		t.Fatalf("expected deterministic output, got %s and %s", uti1, uti2)
	}

	if len(uti1) > utiMaxLength {
		t.Fatalf("UTI exceeds maximum length: %d", len(uti1))
	}

	allowed := regexp.MustCompile(`^[A-Z0-9]+$`)
	if !allowed.MatchString(uti1) {
		t.Fatalf("UTI contains invalid characters: %s", uti1)
	}

	utiNoParts, err := GenerateUTIFromParts("5493004W1IPC50878Z34")
	if err != nil {
		t.Fatalf("unexpected error without parts: %v", err)
	}
	if len(utiNoParts) > utiMaxLength {
		t.Fatalf("UTI exceeds maximum length without parts: %d", len(utiNoParts))
	}
	if !allowed.MatchString(utiNoParts) {
		t.Fatalf("UTI without parts contains invalid characters: %s", utiNoParts)
	}
}
