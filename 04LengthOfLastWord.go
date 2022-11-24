// https://leetcode.com/problems/length-of-last-word/description/
func lengthOfLastWord(s string) int {

	i := len(s) - 1
	length := 0

	for i >= 0 {
		curr := s[i]
		if unicode.IsLetter(rune(curr)) {
			length++
		} else if length > 0 {
			return length
		}
		i--
	}
	return length

}