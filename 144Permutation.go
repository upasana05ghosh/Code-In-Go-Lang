//https://leetcode.com/problems/permutations/description/
func permute(a []int) [][]int {
	var res [][]int
	find(a, &res, []int{}, make([]bool, len(a)))
	return res
}

func find(a []int, res *[][]int, temp []int, used []bool) {
	if len(temp) == len(a) {
		copiedTemp := make([]int, len(temp))
		copy(copiedTemp, temp)
		*res = append(*res, copiedTemp)
		return
	}

	for i := 0; i < len(a); i++ {
		if used[i] {
			continue
		}
		temp = append(temp, a[i])
		used[i] = true
		find(a, res, temp, used)
		used[i] = false
		temp = temp[:len(temp)-1]
	}
}