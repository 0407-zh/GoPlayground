package topinterviewquestions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"ab", "a"}, "a"},
		{[]string{"a"}, "a"},
		{[]string{"abab", "aba", ""}, ""},
		{[]string{"", "b"}, ""},
	}

	for _, test := range tests {
		result := longestCommonPrefix(test.strs)
		assert.Equal(test.expected, result, "Expected \"%s\", but got \"%s\" for input: %v", test.expected, result, test.strs)
	}
}
