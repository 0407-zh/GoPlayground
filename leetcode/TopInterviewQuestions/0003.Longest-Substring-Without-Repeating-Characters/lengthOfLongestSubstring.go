package topinterviewquestions

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)

	right, ans := -1, 0
	for left := 0; left < n; left++ {
		if left != 0 {
			delete(m, s[left-1])
		}

		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		ans = max(ans, right-left+1)
	}

	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
