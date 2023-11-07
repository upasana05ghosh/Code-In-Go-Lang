//https://leetcode.com/problems/climbing-stairs/description/
func climbStairs(n int) int {
	a := make([]int, n+1)

	if n == 1 || n == 2 {
		return n
	}

	a[0] = 1
	a[1] = 1

	for i := 2; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	return a[n]
}