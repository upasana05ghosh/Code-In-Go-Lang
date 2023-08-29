//https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
func maxVowels(s string, k int) int {
	start, end := 0, 0
	count, max := 0, 0

	for end < k && end < len(s) {
		if isVowel(s[end]) {
			count++
		}
		end++
	}

	max = count

	for end < len(s) {
		//remove window
		if isVowel(s[start]) {
			count--
		}
		start++
		//add window
		if isVowel(s[end]) {
			count++
		}
		end++

		max = findMax(max, count)
	}

	return max
}

func findMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func isVowel(a byte) bool {
	if a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' {
		return true
	}
	return false
}