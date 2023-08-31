//https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
func longestSubarray(a []int) int {
	// longest array with atmost 1 zero
	i, j, zero := 0, 0, 0

	for j < len(a) {
		if a[j] == 0 {
			zero++
		}
		if zero > 1 {
			if a[i] == 0 {
				zero--
			}
			i++
		}
		j++
	}

	return j - i - 1
}