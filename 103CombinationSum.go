//https://leetcode.com/problems/combination-sum-iii/description/
func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	temp := []int{}
	backtrack(k, n, &res, &temp, 1)
	return res
}

func backtrack(k int, sum int, res *[][]int, temp *[]int, start int) {
	if sum < 0 {
		return
	}
	if k == 0 && sum == 0 {
		tempCopy := make([]int, len(*temp)) // Create a copy of temp
		copy(tempCopy, *temp)
		*res = append(*res, tempCopy)
		return
	}
	for i := start; i <= 9; i++ {
		*temp = append(*temp, i)
		backtrack(k-1, sum-i, res, temp, i+1)
		*temp = (*temp)[:len(*temp)-1]
	}
}