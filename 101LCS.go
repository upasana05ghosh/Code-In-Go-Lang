//https://leetcode.com/problems/longest-common-subsequence/description/
func longestCommonSubsequence(text1 string, text2 string) int {
	r := len(text1)
	c := len(text2)

	mem := make([][]int, r+1)
	for i := range mem {
		mem[i] = make([]int, c+1)
	}

	//col
	for j := 1; j < c+1; j++ {
		for i := 1; i < r+1; i++ {
			if text1[i-1] == text2[j-1] {
				mem[i][j] = 1 + mem[i-1][j-1]
			} else {
				mem[i][j] = max(mem[i-1][j], mem[i][j-1])
			}
		}
	}
	return mem[r][c]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}