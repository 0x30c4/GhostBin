package utils

import "testing"

func TestParseSizeToBytes(t *testing.T) {
	// Define test cases
	tests := []struct {
		input    string
		expected uint64
		hasError bool
	}{
		// Valid cases
		{"2B", 2, false},
		{"2KB", 2048, false},
		{"2Kb", 256, false},
		{"2MB", 2 * 1024 * 1024, false},
		{"2Mb", 2 * 1024 * 1024 / 8, false},
		{"2GB", 2 * 1024 * 1024 * 1024, false},
		{"2Gb", 2 * 1024 * 1024 * 1024 / 8, false},

		// Edge cases
		{"0B", 0, false},
		{"1B", 1, false},
		{"1KB", 1024, false},

		// Invalid cases
		{"", 0, true},              // Empty input
		{"2", 0, true},             // Missing unit
		{"2XYZ", 0, true},          // Invalid unit
		{"invalid", 0, true},       // Non-numeric input
		{"2.5MB", 0, true},         // Invalid numeric format
		{"-2MB", 0, true},          // Negative numbers
		{"10MB B", 0, true},        // Invalid format with space
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseSizeToBytes(tt.input)

			// Check if we expect an error
			if tt.hasError && err == nil {
				t.Errorf("expected an error for input: %s, got none", tt.input)
			}

			// If no error expected, check the result
			if !tt.hasError {
				if err != nil {
					t.Errorf("did not expect error for input: %s, got: %v", tt.input, err)
				}
				if got != tt.expected {
					t.Errorf("for input: %s, expected %d bytes, got %d bytes", tt.input, tt.expected, got)
				}
			}
		})
	}
}
