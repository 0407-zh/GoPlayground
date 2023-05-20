package topinterviewquestions

func romanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result, currentValue, lastValue := 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		currentValue = symbols[s[i]]

		if currentValue < lastValue {
			result -= currentValue
		} else {
			result += currentValue
		}

		lastValue = currentValue
	}

	return result
}
