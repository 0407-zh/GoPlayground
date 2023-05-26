package topinterviewquestions

import (
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"Case 1: Target is present", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"Case 2: Target is not present", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"Case 3: Array is not rotated", []int{0, 1, 2, 4, 5, 6, 7}, 3, -1},
		{"Case 4: Array is empty", []int{}, 3, -1},
		{"Case 5: Array has one element, target is present", []int{1}, 1, 0},
		{"Case 6: Array has one element, target is not present", []int{1}, 0, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := search(tc.nums, tc.target)
			if got != tc.want {
				t.Errorf("search(%v, %d) = %d; want %d", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
