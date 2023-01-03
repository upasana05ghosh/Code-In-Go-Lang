//https://leetcode.com/problems/search-insert-position/description/
func searchInsert(a []int, target int) int {
	start := 0
	end := len(a) - 1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if a[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return start
}