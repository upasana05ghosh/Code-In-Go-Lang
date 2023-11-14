//https://leetcode.com/problems/word-search/description/
func exist(board [][]byte, word string) bool {
	r := len(board)
	c := len(board[0])

	visit := make([][]bool, r)
	for i := range visit {
		visit[i] = make([]bool, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if present(board, i, j, visit, word, 0, r, c) {
				return true
			}
		}
	}

	return false
}

func present(board [][]byte, i, j int, visit [][]bool, word string, curr, r, c int) bool {
	if curr == len(word) {
		return true
	}

	dist := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	if i >= 0 && j >= 0 && i < r && j < c && board[i][j] == word[curr] && !visit[i][j] {
		visit[i][j] = true

		// Check other directions
		for _, v := range dist {
			x := i + v[0]
			y := j + v[1]

			if present(board, x, y, visit, word, curr+1, r, c) {
				return true
			}
		}

		visit[i][j] = false
	}

	return false
}