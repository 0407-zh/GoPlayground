package topinterviewquestions

func longestPalindrome(s string) string {
	n := len(s)

	if n < 2 {
		return s
	}

	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	maxLength := 1
	start := 0

	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}

			// 如果找到一个回文子串，且它的长度大于之前找到的最长回文子串的长度，
			// 更新最长回文子串的起始和结束位置。
			if dp[i][j] && j-i+1 > maxLength {
				maxLength = j - i + 1
				start = i
			}
		}
	}

	return s[start : start+maxLength]
}
