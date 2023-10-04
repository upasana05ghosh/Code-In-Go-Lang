//https://leetcode.com/problems/counting-bits/description/
func countBits(n int) []int {
	res := make([]int, n+1)

	res[0] = 0
	if n == 0 {
		return res
	}
	res[1] = 1
	if n == 1 {
		return res
	}
	res[2] = 1

	for i := 3; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i-1] + 1
		}
	}

	return res
}
