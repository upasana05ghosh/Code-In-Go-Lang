//https://leetcode.com/problems/unique-paths/description/
func uniquePaths(r int, c int) int {
	if r == 0 {
		return 0
	}

	m := make([][]int, r)
	for i := range m {
		m[i] = make([]int, c)
	}

	//Vertical move
	for i := 0; i < r; i++ {
		m[i][0] = 1
	}

	//Horizontal move
	for i := 0; i < c; i++ {
		m[0][i] = 1
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			m[i][j] = m[i-1][j] + m[i][j-1]
		}
	}

	return m[r-1][c-1]
}