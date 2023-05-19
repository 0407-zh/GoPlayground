package topinterviewquestions

func reverse(x int) int {
	res := 0

	for x != 0 {
		digit := x % 10
		x = x / 10

		if res > 214748364 || (res == 214748364 && digit > 7) {
			return 0
		}

		if res < -214748364 || (res == -214748364 && digit < -8) {
			return 0
		}

		res = res*10 + digit
	}

	return res
}
