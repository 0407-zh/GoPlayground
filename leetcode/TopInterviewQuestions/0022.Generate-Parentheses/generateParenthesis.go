package topinterviewquestions

var strs []string

func generateParenthesis(n int) []string {
	left, right := n, n
	strs = []string{}
	dfs("", left, right)
	return strs
}

func dfs(str string, left, right int) {
	if left == 0 && right == 0 {
		strs = append(strs, str)
		return
	}

	if left > 0 {
		dfs(str+"(", left-1, right)
	}

	if right > 0 && left < right {
		dfs(str+")", left, right-1)
	}
}
