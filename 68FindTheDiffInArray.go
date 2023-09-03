//https://leetcode.com/problems/find-the-difference-of-two-arrays/description/
func findDifference(a []int, b []int) [][]int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	for _, v := range a {
		map1[v] = true
	}

	for _, v := range b {
		map2[v] = true
	}

	res := make([][]int, 2)

	for k, _ := range map1 {
		if _, pres := map2[k]; !pres {
			res[0] = append(res[0], k)
		}
	}

	for k, _ := range map2 {
		if _, pres := map1[k]; !pres {
			res[1] = append(res[1], k)
		}
	}

	return res
}