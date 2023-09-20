//https://leetcode.com/problems/determine-if-two-strings-are-close/description
func closeStrings(w1 string, w2 string) bool {
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)

	for _, v := range w1 {
		freq1[v-'a']++
	}

	for _, v := range w2 {
		freq2[v-'a']++
	}

	//check if unqiue char are present
	for i := 0; i < 26; i++ {
		if (freq1[i] > 0 && freq2[i] == 0) || (freq1[i] == 0 && freq2[i] > 0) {
			return false
		}
	}

	sort.Ints(freq1)
	sort.Ints(freq2)

	// check if freq is same
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true

}