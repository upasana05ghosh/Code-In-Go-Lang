//https://leetcode.com/problems/number-of-provinces/description/
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	isVisited := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if !isVisited[i] {
			dfs(isConnected, isVisited, i, n)
			count++
		}
	}

	return count
}

func dfs(isConnected [][]int, isVisited []bool, i, n int) {
	isVisited[i] = true

	for j := 0; j < n; j++ {
		if isConnected[i][j] == 1 && !isVisited[j] {
			dfs(isConnected, isVisited, j, n)
		}
	}
}

