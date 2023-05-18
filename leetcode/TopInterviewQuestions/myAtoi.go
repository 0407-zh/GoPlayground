package topinterviewquestions

import "math"

func myAtoi(s string) int {
	var digits = map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	res, sign, i := 0, 1, 0

	// Ignore leading whitespaces
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}

	// Check for sign
	if i < len(s) {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}

	// Process numeric part of the string
	for i < len(s) {
		if num, ok := digits[s[i]]; ok {
			// Check if result will be out of bounds, i.e., less than INT_MIN or greater than INT_MAX
			if res > (math.MaxInt32-num)/10 {
				if sign == -1 {
					return math.MinInt32
				}
				return math.MaxInt32
			}

			res = res*10 + num
			i++
		} else {
			// Break on encountering a non-digit character
			break
		}
	}

	return sign * res
}
