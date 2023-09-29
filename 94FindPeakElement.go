//https://leetcode.com/problems/find-peak-element/description/
func findPeakElement(a []int) int {
	n := len(a)
	l, r := 0, n-1

	for l <= r {
		mid := (l + r) / 2

		if (mid == 0 || a[mid] > a[mid-1]) && (mid == n-1 || a[mid] > a[mid+1]) {
			return mid
		}

		if a[mid] < a[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}