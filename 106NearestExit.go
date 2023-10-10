//https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/description/
func nearestExit(maze [][]byte, entrance []int) int {
	var q [][2]int
	steps := 0
	r, c := len(maze), len(maze[0])

	//update entrance so we don't step into it again
	maze[entrance[0]][entrance[1]] = '+'
	q = append(q, [2]int{entrance[0], entrance[1]})

	dist := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(q) > 0 {
		size := len(q)
		steps++

		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:] //pop queue

			for _, v := range dist {
				x := v[0] + n[0]
				y := v[1] + n[1]

				if x >= 0 && y >= 0 && x < r && y < c && maze[x][y] == '.' {
					if x == 0 || y == 0 || x == r-1 || y == c-1 {
						return steps
					} else {
						maze[x][y] = '+'
						q = append(q, [2]int{x, y})
					}
				}
			}
		}
	}
	return -1
}