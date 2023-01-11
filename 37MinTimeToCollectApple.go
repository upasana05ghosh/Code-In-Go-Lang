func minTime(n int, edges [][]int, hasApple []bool) int {
	graph := make(map[int][]int)
	visit := make(map[int]bool)
	buildGraph(edges, graph)
	return calTime(0, graph, visit, hasApple)
}

func buildGraph(edges [][]int, graph map[int][]int) {
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if _, ok := graph[x]; !ok {
			graph[x] = []int{}
		}
		graph[x] = append(graph[x], y)
		if _, ok := graph[y]; !ok {
			graph[y] = []int{}
		}
		graph[y] = append(graph[y], x)
	}
}

func calTime(node int, graph map[int][]int, visit map[int]bool, hasApple []bool) int {
	visit[node] = true

	res := 0

	for _, val := range graph[node] {
		if !visit[val] {
			res += calTime(val, graph, visit, hasApple)
		}
	}

	if node == 0 {
		return res
	}

	if res > 0 || hasApple[node] {
		res += 2
	}

	return res
}