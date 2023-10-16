//https://leetcode.com/problems/successful-pairs-of-spells-and-potions/description/
func successfulPairs(spells []int, potions []int, success int64) []int {
	n := len(spells)
	res := make([]int, n)

	//sort potions
	sort.Ints(potions)

	for i := 0; i < n; i++ {
		j := 0
		for j < len(potions) {
			if int64(spells[i]*potions[j]) >= success {
				break
			}
			j++
		}
		res[i] = len(potions) - j
	}

	return res
}