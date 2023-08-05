//https://leetcode.com/problems/reverse-words-in-a-string/description
func reverseWords(s string) string {
	words := strings.Fields(s) //remove white spaces
	reverse(words)
	return strings.Join(words, " ")
}

func reverse(a []string) {
	i, j := 0, len(a)-1

	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}