//https://leetcode.com/problems/equal-row-and-column-pairs/description/
func equalPairs(grid [][]int) int {
	m := make(map[string]int)
	n := len(grid)

	for i := 0; i < n; i++ {
		var s strings.Builder
		for j := 0; j < n; j++ {
			s.WriteString(strconv.Itoa(grid[i][j]) + "_")
		}
		m[s.String()] += 1
	}

	count := 0

	for j := 0; j < n; j++ {
		var s strings.Builder
		for i := 0; i < n; i++ {
			s.WriteString(strconv.Itoa(grid[i][j]) + "_")
		}
		count += m[s.String()]
	}

	return count
}