//https://leetcode.com/problems/search-in-rotated-sorted-array/description/
func search(a []int, target int) int {
	n := len(a)

	//find the pivot point
	p := findPivot(a, 0, n-1)

	//use normal binary search
	if a[0] <= target && target <= a[p] {
		return binarySearch(a, 0, p, target)
	}

	return binarySearch(a, p+1, n-1, target)
}

func findPivot(a []int, start, end int) int {
	if start <= end {
		mid := (start + end) / 2

		if mid+1 == len(a) || a[mid] > a[mid+1] {
			return mid
		}

		if a[start] <= a[mid] { //search in next half
			return findPivot(a, mid+1, end)
		}

		return findPivot(a, start, mid-1)
	}
	return -1
}

func binarySearch(a []int, start, end, target int) int {
	for start <= end {
		mid := (start + end) / 2

		if a[mid] == target {
			return mid
		}

		if a[mid] > target { //search in left half
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}