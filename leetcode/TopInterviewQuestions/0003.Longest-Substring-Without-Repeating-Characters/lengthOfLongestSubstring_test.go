package topinterviewquestions

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abba", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.input)
			if result != tc.output {
				t.Errorf("Expected %d, but got %d for input %s", tc.output, result, tc.input)
			}
		})
	}
}
