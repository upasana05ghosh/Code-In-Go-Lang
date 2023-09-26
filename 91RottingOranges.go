//https://leetcode.com/problems/rotting-oranges/description
func orangesRotting(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	var q [][2]int
	fresh := 0

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	time := 0
	dir := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for len(q) > 0 {
		time++
		size := len(q)

		for i := 0; i < size; i++ {
			for _, v := range dir {
				x := q[0][0] + v[0]
				y := q[0][1] + v[1]
				if x >= 0 && y >= 0 && x < r && y < c && grid[x][y] == 1 {
					grid[x][y] = 2
					q = append(q, [2]int{x, y})
					fresh--
				}
			}
			q = q[1:]
		}

		if fresh == 0 {
			return time
		}
	}

	return -1
}