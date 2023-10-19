//https://leetcode.com/problems/domino-and-tromino-tiling/description/
func numTilings(n int) int {
	if n == 1 {
		return 1
	}

	first := make([]int, n+1)
	next := make([]int, n+1)

	first[1], first[2] = 1, 2
	next[2] = 1

	for i := 3; i <= n; i++ {
		first[i] = (first[i-2] + first[i-1] + next[i-1]*2) % 1000000007
		next[i] = (first[i-2] + next[i-1]) % 1000000007
	}

	return first[n]
}