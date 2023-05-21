package topinterviewquestions

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums: []int{},
			want: [][]int{},
		},
		{
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}
	for _, test := range tests {
		got := threeSum(test.nums)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("threeSum(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}
