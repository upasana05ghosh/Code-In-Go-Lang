//https://leetcode.com/problems/kth-largest-element-in-an-array/description/
func findKthLargest(a []int, k int) int {
	n := len(a)
	k = n - k
	l, r := 0, n-1

	for l <= r {
		j := partition(a, l, r)
		if j == k {
			break
		}
		if j > k {
			r = j - 1
		} else {
			l = j + 1
		}
	}

	return a[k]
}

func partition(a []int, start, end int) int {
	v, i, j := a[start], start, end

	for i < j {
		for i < len(a) && a[i] <= v {
			i++
		}
		for a[j] > v {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[start], a[j] = a[j], a[start]
	return j
}