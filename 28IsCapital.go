//https://leetcode.com/problems/detect-capital/description/
func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}

	if !isCapital(rune(word[0])) && isCapital(rune(word[1])) {
		return false
	}

	isCaps := false
	if isCapital(rune(word[1])) { //if second is caps
		isCaps = true
	}

	n := len(word)
	i := 2
	for i < n {
		if isCaps != isCapital(rune(word[i])) {
			return false
		}
		i++
	}

	return true
}

func isCapital(ch rune) bool {
	if ch >= 'A' && ch <= 'Z' { //if second is caps
		return true
	}
	return false
}