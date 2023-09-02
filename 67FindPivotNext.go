// /https://leetcode.com/problems/find-pivot-index/description/
func pivotIndex(a []int) int {
	r, l := 0, 0

	for _, v := range a {
		r += v
	}

	for i, v := range a {
		r -= v
		if i != 0 {
			l += a[i-1]
		}
		if l == r {
			return i
		}
	}
	return -1
}