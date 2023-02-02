//https://leetcode.com/problems/verifying-an-alien-dictionary/description/
func isAlienSorted(words []string, order string) bool {
	orderMap := make([]int, len(order))

	for i, w := range order {
		orderMap[w-'a'] = i
	}

	for i := 1; i < len(words); i++ {
		if unSorted(words[i-1], words[i], orderMap) {
			return false
		}
	}

	return true
}

func unSorted(w1 string, w2 string, orderMap []int) bool {
	n1, n2 := len(w1), len(w2)

	for i := 0; i < n1 && i < n2; i++ {
		if rune(w1[i]) != rune(w2[i]) {
			return orderMap[w1[i]-'a'] > orderMap[w2[i]-'a']
		}
	}

	return n1 > n2
}
