//https://leetcode.com/problems/minimum-path-sum/description/
func minPathSum(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	mem := make([][]int, r)
	for i := range mem {
		mem[i] = make([]int, c)
	}
	return findMinPath(grid, 0, 0, 0, r, c, mem)
}

func findMinPath(grid [][]int, i, j, sum, r, c int, mem [][]int) int {
	if i == r-1 && j == c-1 { //found end
		return grid[i][j] + sum
	}

	//invalid
	if i >= r || j >= c {
		return math.MaxInt32 //return max
	}

	if mem[i][j] != 0 {
		return mem[i][j] //already computed
	}

	right := findMinPath(grid, i, j+1, sum, r, c, mem)
	bottom := findMinPath(grid, i+1, j, sum, r, c, mem)

	mem[i][j] = grid[i][j] + min(right, bottom)
	return mem[i][j]
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
} 