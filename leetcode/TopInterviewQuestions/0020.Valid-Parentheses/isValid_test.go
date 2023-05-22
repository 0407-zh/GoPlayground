package topinterviewquestions

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Example 1",
			input:    "()",
			expected: true,
		},
		{
			name:     "Example 2",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "Example 3",
			input:    "(]",
			expected: false,
		},
		{
			name:     "Example 4",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "Example 5",
			input:    "{[]}",
			expected: true,
		},
		{
			name:     "Example 6",
			input:    "]",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.input)

			if result != tt.expected {
				t.Errorf("Test case '%s' failed: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}
