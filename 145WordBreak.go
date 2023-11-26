//https://leetcode.com/problems/word-break/description/
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	return findBreak(s, wordSet, make(map[string]bool))
}

func findBreak(s string, wordSet map[string]bool, present map[string]bool) bool {
	if v, p := present[s]; p {
		return v
	}

	if _, p := wordSet[s]; p {
		return true
	}

	for i := 1; i < len(s); i++ {
		prefix := s[:i]
		if _, p := wordSet[prefix]; p {
			if findBreak(s[i:], wordSet, present) {
				present[s] = true
				return true
			}
		}
	}

	present[s] = false
	return false
}


