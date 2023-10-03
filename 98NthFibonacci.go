//https://leetcode.com/problems/n-th-tribonacci-number/description/
func tribonacci(n int) int {
	a := make([]int, n+1)

	if n == 0 || n == 1 {
		return n
	}

	if n == 2 {
		return 1
	}

	a[0] = 0
	a[1] = 1
	a[2] = 1

	for i := 3; i <= n; i++ {
		a[i] = a[i-1] + a[i-2] + a[i-3]
	}

	return a[n]
}