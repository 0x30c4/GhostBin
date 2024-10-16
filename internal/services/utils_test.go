package services

import (
	"testing"
	"regexp"
	"strings"
)

// TestRandomPasteIdPrefix tests the randomPasteIdPrefix function
func TestRandomPasteIdPrefix(t *testing.T) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetPattern := "^[a-zA-Z0-9]+$"
	regex := regexp.MustCompile(charsetPattern)

	// Define test cases for different lengths
	tests := []struct {
		length uint8
	}{
		{5},
		{10},
		{20},
		{50},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run("LengthTest", func(t *testing.T) {
			result := randomPasteIdPrefix(tt.length)

			// Check if the result is of the correct length
			if len(result) != int(tt.length) {
				t.Errorf("expected length: %d, got: %d", tt.length, len(result))
			}

			// Check if the result contains only valid characters from the charset
			if !regex.MatchString(result) {
				t.Errorf("generated string contains invalid characters: %s", result)
			}

			// Check if all characters in the result are from the charset
			for _, ch := range result {
				if !strings.ContainsRune(charset, ch) {
					t.Errorf("unexpected character: %c", ch)
				}
			}
		})
	}
}

// BenchmarkRandomPasteIdPrefix benchmarks the randomPasteIdPrefix function
func BenchmarkRandomPasteIdPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randomPasteIdPrefix(20)
	}
}
