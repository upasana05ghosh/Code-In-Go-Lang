//https://leetcode.com/problems/find-closest-node-to-given-two-nodes/description/
func closestMeetingNode(edges []int, node1 int, node2 int) int {
    //find min from both src shortest path
    n := len(edges)

    dist1 := getArrayWithVal(n, math.MaxInt)
    dist2 := getArrayWithVal(n, math.MaxInt)

    bfs(edges, &dist1, node1)
    bfs(edges, &dist2, node2)

    resp := math.MaxInt
    respNode := -1
    for i, v:= range dist1 {
        if v == math.MaxInt || dist2[i] == math.MaxInt {
            continue
        }
        if max(v, dist2[i]) < resp {
            resp = max(v, dist2[i])
            respNode = i
        }
    }


    return respNode
}

func max(x, y int) int {
    if x >= y {
        return x
    }
    return y
}

func bfs(edges []int, dist *[]int, src int) {
    q:= []int{}
    q = append(q, src)
    (*dist)[src] = 0
    for len(q) > 0 {
        n := len(q)
        v := q[0]
        q = append(q[1:n]) //queue poll
        if edges[v] != -1 && (*dist)[edges[v]] == math.MaxInt {
            q = append(q, edges[v])
            (*dist)[edges[v]] = (*dist)[v] + 1
        }
    }
}

func getArrayWithVal(n int, v int) [] int {
    dist := make([]int, n)

    for i:= range dist {
        dist[i] = v
    }
    return dist
}