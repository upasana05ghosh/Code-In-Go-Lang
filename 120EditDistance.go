// /https://leetcode.com/problems/edit-distance/description/
func minDistance(w1 string, w2 string) int {
	r, c := len(w1), len(w2)
	if r == 0 {
		return c
	}

	if c == 0 {
		return r
	}

	a := make([][]int, r)
	for i := range a {
		a[i] = make([]int, c)
	}

	return match(w1, w2, 0, 0, a)
}

func match(w1 string, w2 string, i, j int, a [][]int) int {
	if len(w1) == i {
		return len(w2) - j
	}

	if len(w2) == j {
		return len(w1) - i
	}

	if a[i][j] != 0 {
		return a[i][j]
	}

	if w1[i] == w2[j] {
		return match(w1, w2, i+1, j+1, a)
	}

	insert := match(w1, w2, i+1, j, a)
	delete := match(w1, w2, i, j+1, a)
	replace := match(w1, w2, i+1, j+1, a)

	a[i][j] = min(insert, min(delete, replace)) + 1
	return a[i][j]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}