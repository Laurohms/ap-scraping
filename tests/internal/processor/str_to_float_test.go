package processor_test

import (
	"testing"

	"github.com/Laurohms/ap-scraper/internal/processors"
)

func TestStrToFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
		hasError bool
	}{
		{"Positive number with +", " + 1.234,05", 1234.05, false},
		{"Positive number without +", "1.234,05", 1234.05, false},
		{"Negative number", "-1.234,05", -1234.05, false},
		{"Number with spaces", "   1.234,05   ", 1234.05, false},
		{"Number without thousand separator", "1234,05", 1234.05, false},
		{"Empty string", "", 0, true},
		{"Invalid format with letters", "abc", 0, true},
		{"Invalid format with mixed characters", "12.34.56", 123456.00, false},
		{"Negative number with spaces", "   - 1.234,05   ", -1234.05, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processors.StrToFloat(tt.input)

			if (err != nil) != tt.hasError {
				t.Errorf("StrToFloat(%q) error = %v, expected error = %v", tt.input, err, tt.hasError)
				return
			}

			if got != tt.expected {
				t.Errorf("StrToFloat(%q) = %f, want %f", tt.input, got, tt.expected)
			}
		})
	}
}
