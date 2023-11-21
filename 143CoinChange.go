//https://leetcode.com/problems/coin-change/description/
func coinChange(coins []int, amt int) int {
	a := make([]int, amt+1)

	for i := 0; i <= amt; i++ {
		a[i] = amt + 1
	}

	a[0] = 0

	for _, v := range coins {
		for i := v; i <= amt; i++ {
			a[i] = min(a[i], a[i-v]+1)
		}
	}

	if a[amt] == amt+1 {
		return -1
	}

	return a[amt]
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}