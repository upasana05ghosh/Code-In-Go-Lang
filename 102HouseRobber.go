//https://leetcode.com/problems/house-robber/description/

func rob(a []int) int {
	curr := a[0]
	prev := 0

	for i := 1; i < len(a); i++ {
		next := prev + a[i]
		prev = max(curr, prev)
		curr = next
	}

	return max(curr, prev)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}