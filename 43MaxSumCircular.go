//https://leetcode.com/problems/maximum-sum-circular-subarray/description/
func maxSubarraySumCircular(nums []int) int {
	//max = maxInStraight, total - minInStraight

	//cal total
	total := 0
	for _, v := range nums {
		total += v
	}
	//cal maxSum using Kadane's algo
	maxSum, sum := math.MinInt, 0
	for _, v := range nums {
		sum += v
		maxSum = max(maxSum, sum)
		if sum < 0 {
			sum = 0
		}
	}
	//cal minSum
	minSum, s := math.MaxInt, 0
	for _, v := range nums {
		s += v
		minSum = min(minSum, s)
		if s > 0 {
			s = 0
		}
	}
	if maxSum < 0 {
		return maxSum
	}
	return max(maxSum, total-minSum)
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}