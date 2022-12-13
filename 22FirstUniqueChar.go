//https://leetcode.com/problems/first-unique-character-in-a-string/description/

func firstUniqChar(s string) int {
	var a [26]int

	for _, c := range s {
		a[rune(c)-'a']++
	}

	for i, c := range s {
		if a[rune(c)-'a'] == 1 {
			return i
		}
	}

	return -1
}
