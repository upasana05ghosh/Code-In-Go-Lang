//https://github.com/upasana05ghosh/Code-In-Go-Lang/blob/main/05ClimbingStirs.go
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	a := make([]int, n+1)
	a[0] = 1
	a[1] = 1

	for i := 2; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	return a[n]
}
