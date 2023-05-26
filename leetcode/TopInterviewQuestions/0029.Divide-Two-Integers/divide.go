package topinterviewquestions

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if dividend == math.MaxInt32 && divisor == -1 {
		return math.MinInt32 + 1
	}

	a, b := abs(dividend), abs(divisor)
	sign := 1

	if dividend < 0 {
		sign *= -1
	}
	if divisor < 0 {
		sign *= -1
	}

	res := 0
	for a >= b {
		temp := b
		multiple := 1
		if a >= temp<<1 {
			temp <<= 1
			multiple <<= 1
		}

		a -= temp
		res += multiple
	}

	return res * sign
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
