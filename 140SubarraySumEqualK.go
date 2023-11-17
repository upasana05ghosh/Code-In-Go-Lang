//https://leetcode.com/problems/subarray-sum-equals-k/description/
func subarraySum(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	m[0] = 1
	sum := 0

	for _, v := range nums {
		sum += v
		if val, isPresent := m[sum-k]; isPresent {
			res += val
		}
		m[sum] += 1
	}

	return res
}