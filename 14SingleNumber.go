//https://leetcode.com/problems/single-number/description/
func singleNumber(nums []int) int {
	x := 0

	for _, val := range nums {
		x ^= val
	}

	return x
}