//https://leetcode.com/problems/word-pattern/description/
func wordPattern(pattern string, s string) bool {

	patternMap := make(map[rune]string)
	patternExist := make(map[string]bool)

	ss := strings.Split(s, " ")

	if len(pattern) != len(ss) {
		return false
	}

	for i, val := range pattern {
		pp, ok := patternMap[val]
		if ok {
			if pp != ss[i] {
				return false
			}
		} else {
			_, ok := patternExist[ss[i]] //check if pattern not existing for others
			if ok {
				return false
			}
			patternExist[ss[i]] = true
		}

		patternMap[val] = ss[i]
	}

	return true

}