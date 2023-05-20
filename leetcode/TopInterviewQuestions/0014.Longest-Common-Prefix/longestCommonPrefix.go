package topinterviewquestions

func longestCommonPrefix(strs []string) string {
	p := strs[0]

	for _, str := range strs {
		i := 0
		for ; i < len(p) && i < len(str) && p[i] == str[i]; i++ {

		}
		p = p[:i]
	}

	return p
}
