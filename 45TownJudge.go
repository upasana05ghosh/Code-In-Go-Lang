//https://leetcode.com/problems/find-the-town-judge/description/
func findJudge(n int, trust [][]int) int {
	tt := make([]int, n+1)

	for _, v := range trust {
		tt[v[0]]--
		tt[v[1]]++
	}

	for i := 1; i <= n; i++ {
		if tt[i] == n-1 {
			return i
		}
	}
	return -1
}