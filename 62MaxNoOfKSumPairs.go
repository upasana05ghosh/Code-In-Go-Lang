//https://leetcode.com/problems/max-number-of-k-sum-pairs/description
func maxOperations(a []int, k int) int {
	sumMap := make(map[int]int)
	count := 0

	for _, v := range a {
		diff := k - v
		if val, isPresent := sumMap[diff]; isPresent && val > 0 {
			count++
			sumMap[diff] = val - 1
		} else {
			sumMap[v] = sumMap[v] + 1
		}
	}
	return count
}