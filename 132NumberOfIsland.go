//https://leetcode.com/problems/number-of-islands/description/
func numIslands(grid [][]byte) int {
	r := len(grid)
	c := len(grid[0])

	visit := make([][]bool, r)
	for i := range visit {
		visit[i] = make([]bool, c)
	}
	count := 0

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' && !visit[i][j] {
				dfs(grid, i, j, r, c, visit)
				count++
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i int, j int, r int, c int, visit [][]bool) {
	visit[i][j] = true

	dist := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for _, d := range dist {
		x := i + d[0]
		y := j + d[1]

		if x >= 0 && y >= 0 && x < r && y < c && !visit[x][y] && grid[x][y] == '1' {
			dfs(grid, x, y, r, c, visit)
		}
	}
}

