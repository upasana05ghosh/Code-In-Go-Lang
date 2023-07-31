//https://leetcode.com/problems/merge-strings-alternately/description
func mergeAlternately(word1 string, word2 string) string {
	var comb []rune

	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		comb = append(comb, rune(word1[i]))
		comb = append(comb, rune(word2[j]))
		i++
		j++
	}

	for i < len(word1) {
		comb = append(comb, rune(word1[i]))
		i++
	}

	for j < len(word2) {
		comb = append(comb, rune(word2[j]))
		j++
	}

	return string(comb)
}