//https://leetcode.com/problems/valid-palindrome/description/
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	n := len(s)
	i := 0
	j := n - 1

	for i < j {
		for i < n && isEmpty(rune(s[i])) {
			i++
		}
		for j >= 0 && isEmpty(rune(s[j])) {
			j--
		}

		if i < n && j >= 0 && s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true

}

func isEmpty(c rune) bool {
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
		return false
	}
	return true
}