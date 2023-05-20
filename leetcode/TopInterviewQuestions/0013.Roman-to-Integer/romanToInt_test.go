package topinterviewquestions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		roman    string
		expected int
	}{
		// {"I", 1},
		// {"V", 5},
		// {"X", 10},
		// {"L", 50},
		// {"C", 100},
		// {"D", 500},
		// {"M", 1000},
		// {"IV", 4},
		// {"IX", 9},
		// {"XL", 40},
		// {"XC", 90},
		// {"CD", 400},
		// {"CM", 900},
		// {"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		result := romanToInt(test.roman)
		assert.Equal(test.expected, result, "Expected %d, but got %d for Roman: %s", test.expected, result, test.roman)
	}
}
