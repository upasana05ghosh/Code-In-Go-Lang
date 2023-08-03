//https://leetcode.com/problems/reverse-vowels-of-a-string/description/
func reverseVowels(s string) string {
	var vowel []rune

	for _, v := range s {
		if isVowel(v) {
			vowel = append(vowel, v)
		}
	}
	ss := []rune(s)

	for i, v := range s {
		if isVowel(v) {
			ss[i] = vowel[len(vowel)-1]
			vowel = vowel[:len(vowel)-1]
		}
	}
	return string(ss)
}

func isVowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}