//https://leetcode.com/problems/non-decreasing-subsequences/description/
func findSubsequences(a []int) [][]int {
	res := [][]int{}
	backtrack(&res, a, 0, []int{})
	return res
}

func backtrack(res *[][]int, a []int, index int, curr []int) {
	if len(curr) > 1 {
		*res = append(*res, append([]int{}, curr...)) //slice are passed by reference, so created empty slice and added curr to it :)
	}

	visit := map[int]bool{}

	for i := index; i < len(a); i++ {
		if visit[a[i]] {
			continue
		}
		if len(curr) == 0 || curr[len(curr)-1] <= a[i] {
			visit[a[i]] = true
			curr = append(curr, a[i])
			backtrack(res, a, i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}
}