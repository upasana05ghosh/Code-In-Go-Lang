//https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/description/
func countSubTrees(n int, edges [][]int, labels string) []int {
	graph := make(map[int][]int)
	buildGraph(edges, graph)
	res := make([]int, n)
	visit := make(map[int]bool)
	buildRes(0, res, graph, labels, visit)
	return res
}

func buildGraph(edges [][]int, graph map[int][]int) {
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if _, ok := graph[a]; !ok {
			graph[a] = []int{}
		}
		if _, ok := graph[b]; !ok {
			graph[b] = []int{}
		}
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
}

func buildRes(curr int, res []int, graph map[int][]int, labels string, visit map[int]bool) map[rune]int {
	visit[curr] = true
	occMap := make(map[rune]int)

	for _, child := range graph[curr] {
		if !visit[child] {
			childMap := buildRes(child, res, graph, labels, visit)
			//update occMap
			for k, v := range childMap {
				occMap[k] += v
			}
		}
	}

	occMap[rune(labels[curr])] += 1
	res[curr] = occMap[rune(labels[curr])]
	return occMap
}