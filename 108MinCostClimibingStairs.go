//https://leetcode.com/problems/min-cost-climbing-stairs/description
func minCostClimbingStairs(cost []int) int {
	prev := cost[0]
	next := cost[1]

	for i := 2; i < len(cost); i++ {
		//update next with min
		if prev <= next {
			prev, next = next, prev+cost[i]
		} else {
			prev, next = next, next+cost[i]
		}
	}

	if prev <= next {
		return prev
	}
	return next
}