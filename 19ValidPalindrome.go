//https://leetcode.com/problems/valid-anagram/description/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, val := range s {
		map1[val]++
	}

	for _, val := range t {
		map2[val]++
	}

	for key, val := range map1 {
		if val != map2[key] {
			return false
		}
	}

	return true
}