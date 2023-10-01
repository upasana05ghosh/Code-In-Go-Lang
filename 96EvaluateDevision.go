//https://leetcode.com/problems/evaluate-division/description/
func calcEquation(e [][]string, values []float64, qq [][]string) []float64 {
	vertexMap := make(map[string][]Vertex)

	//build graph
	for i, v := range e {
		a1 := v[0]
		a2 := v[1]
		vertexMap[a1] = append(vertexMap[a1], Vertex{a2, values[i]})
		vertexMap[a2] = append(vertexMap[a2], Vertex{a1, 1.0 / values[i]})
	}

	res := make([]float64, len(qq))

	for i, q := range qq {
		a1 := q[0]
		a2 := q[1]
		visited := make(map[string]bool)

		//key not present in map
		if _, ok := vertexMap[a1]; !ok {
			res[i] = -1.0
		} else {
			//dfs
			res[i], _ = dfs(a1, a2, vertexMap, visited)
		}
	}
	return res
}

func dfs(start, end string, vertexMap map[string][]Vertex, visited map[string]bool) (float64, bool) {
	// if same
	if start == end {
		return 1.0, true
	}

	for _, v := range vertexMap[start] {
		if visited[v.Name] {
			continue
		}
		visited[v.Name] = true
		if val, ok := dfs(v.Name, end, vertexMap, visited); ok {
			return val * v.Val, true
		}
	}
	return -1.0, false
}

type Vertex struct {
	Name string
	Val  float64
}