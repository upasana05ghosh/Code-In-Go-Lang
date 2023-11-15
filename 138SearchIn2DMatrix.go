//https://leetcode.com/problems/search-a-2d-matrix/description/
func searchMatrix(matrix [][]int, target int) bool {
	r, c := len(matrix), len(matrix[0])

	i, j := 0, c-1

	for i < r && j >= 0 {
		fmt.Println(matrix[i][j])
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}

	return false
}