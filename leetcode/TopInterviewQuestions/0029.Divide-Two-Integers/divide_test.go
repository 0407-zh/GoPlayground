package topinterviewquestions

import (
	"testing"
	"time"
)

func TestDivide(t *testing.T) {
	testCases := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{"Case 1: Positive integers", 10, 3, 3},
		{"Case 2: Negative dividend", -7, 3, -2},
		{"Case 3: Negative divisor", 7, -3, -2},
		{"Case 4: Both negative", -7, -3, 2},
		{"Case 5: Divisor is 1", 7, 1, 7},
		{"Case 6: Divisor is -1, no overflow", -7, -1, 7},
		{"Case 7: Divisor is -1, overflow", -2147483648, -1, 2147483647},
		{"Case 8: Dividend is 0", 0, 1, 0},
		{"Case 9: Large dividend and divisor", 2147483647, -1, -2147483647},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := make(chan int, 1)
			go func() {
				c <- divide(tc.dividend, tc.divisor)
			}()

			select {
			case res := <-c:
				if res != tc.want {
					t.Errorf("divide(%d, %d) = %d; want %d", tc.dividend, tc.divisor, res, tc.want)
				}
			case <-time.After(1 * time.Second):
				t.Errorf("divide(%d, %d) timed out", tc.dividend, tc.divisor)
			}
		})
	}
}
