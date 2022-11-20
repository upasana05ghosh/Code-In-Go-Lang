func twoSum(arr []int, target int) []int {
	targetMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		v := target - arr[i]
		j, ok := targetMap[arr[i]]
		if ok {
			return []int{i, j}
		}
		targetMap[v] = i
	}

	return []int{}
}