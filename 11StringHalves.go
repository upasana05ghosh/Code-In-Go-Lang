//https://leetcode.com/problems/determine-if-string-halves-are-alike/description/
func halvesAreAlike(s string) bool {
	s = strings.ToLower(s)
	f := countVowel(s[:len(s)/2])
	r := countVowel(s[len(s)/2:])
	return f == r
}

func countVowel(s string) int {
	count := 0

	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			count++
		}
	}

	return count
}