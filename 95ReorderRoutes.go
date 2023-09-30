//https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/description/
func minReorder(n int, connections [][]int) int {
	conMap := make(map[int][]int)

	for _, conn := range connections {
		conMap[conn[0]] = append(conMap[conn[0]], conn[1])
		conMap[conn[1]] = append(conMap[conn[1]], -conn[0])
	}

	res := 0
	visited := make([]bool, n)

	dfs(0, &res, visited, conMap)
	return res
}

func dfs(i int, res *int, visited []bool, conMap map[int][]int) {
	visited[i] = true

	for _, v := range conMap[i] {
		if v < 0 {
			v = -v
			if !visited[v] {
				dfs(v, res, visited, conMap)
			}
		} else {
			if !visited[v] {
				*res = *res + 1
				dfs(v, res, visited, conMap)
			}
		}
	}
}