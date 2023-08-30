//https://leetcode.com/problems/max-consecutive-ones-iii/description/
func longestOnes(a []int, k int) int {
	i, j, zero := 0, 0, 0
	for j < len(a) {
		if a[j] == 0 {
			zero++
		}

		if zero > k {
			if a[i] == 0 {
				zero--
			}
			i++
		}
		j++
	}
	return j - i
}