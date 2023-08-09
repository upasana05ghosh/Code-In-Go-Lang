//https://leetcode.com/problems/is-subsequence/description/

func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}

	i := 0
	for j := 0; j < len(t) && i < len(s); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}