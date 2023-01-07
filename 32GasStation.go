//https://leetcode.com/problems/gas-station/description/
func canCompleteCircuit(gas []int, cost []int) int {
	//similar to Kadane's algo

	sum, res, total := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		sum += (gas[i] - cost[i])
		if sum < 0 {
			total += sum
			sum = 0
			res = i + 1
		}
	}

	total += sum

	if total < 0 {
		return -1
	}
	return res
}