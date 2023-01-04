//https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/description/
func minimumRounds(tasks []int) int {
	taskMap := make(map[int]int)
	count := 0

	for _, v := range tasks {
		taskMap[v] += 1
	}

	for _, v := range taskMap {
		if v == 1 {
			return -1
		}
		count += (v + 2) / 3

	}
	return count
}