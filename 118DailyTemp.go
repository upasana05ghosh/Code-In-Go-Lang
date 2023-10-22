//https://leetcode.com/problems/daily-temperatures/description/
func dailyTemperatures(temp []int) []int {
	n := len(temp)
	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		j := i + 1
		noFound := false

		for j < n && temp[i] >= temp[j] {
			if res[j] == 0 {
				noFound = true
				break
			} else {
				j += res[j]
			}
		}

		if j == n || noFound {
			res[i] = 0
		} else {
			res[i] = j - i
		}
	}

	return res
}