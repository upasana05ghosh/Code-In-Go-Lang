//https://leetcode.com/problems/unique-number-of-occurrences/description/
func uniqueOccurrences(a []int) bool {
	m := make(map[int]int)

	for _, v := range a {
		m[v] += 1
	}

	s := make(map[int]bool)

	for _, v := range m {
		if s[v] == true {
			return false
		}
		s[v] = true
	}

	return true
}