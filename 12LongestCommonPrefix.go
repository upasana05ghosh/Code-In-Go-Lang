//https://leetcode.com/problems/longest-common-prefix/description/
func longestCommonPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}

	i := 0
	for i < len(s[0]) {
		c := s[0][i]
		isCharSame := true
		for j := 0; j < len(s); j++ {
			if len(s[j]) <= i || s[j][i] != c {
				isCharSame = false
				break
			}
		}
		if !isCharSame {
			break
		}
		i++
	}

	return s[0][:i]
}