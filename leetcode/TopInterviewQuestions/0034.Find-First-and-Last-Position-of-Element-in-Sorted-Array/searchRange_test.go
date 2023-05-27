package topinterviewquestions

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			nums:   []int{1},
			target: 1,
			want:   []int{0, 0},
		},
		{
			nums:   []int{1},
			target: 0,
			want:   []int{-1, -1},
		},
	}

	for _, testCase := range testCases {
		got := searchRange(testCase.nums, testCase.target)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("Error: expected %v, but got %v", testCase.want, got)
		}
	}
}
