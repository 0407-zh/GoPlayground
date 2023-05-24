package topinterviewquestions

func strStr(haystack string, needle string) int {
	n := len(needle)

	for i := 0; i < len(haystack); i++ {
		if len(haystack) >= i+n {
			subStr := haystack[i : i+n]
			if subStr == needle {
				return i
			}
		} else {
			return -1
		}
	}

	return -1
}
