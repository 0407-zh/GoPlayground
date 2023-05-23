package topinterviewquestions

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	testCases := []struct {
		n        int
		expected []string
	}{
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		// 这里可以添加更多的测试用例
	}

	for _, tc := range testCases {
		result := generateParenthesis(tc.n)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("For n = %d, expected %v, got %v", tc.n, tc.expected, result)
		}
	}
}
