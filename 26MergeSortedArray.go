//https://leetcode.com/problems/merge-sorted-array/description/
func merge(a1 []int, m int, a2 []int, n int) {
	res := make([]int, m+n)
	i, j, k := 0, 0, 0

	for i < m && j < n {
		if a1[i] <= a2[j] {
			res[k] = a1[i]
			i += 1
		} else {
			res[k] = a2[j]
			j += 1
		}
		k += 1
	}

	for i < m && a1[i] != 0 {
		res[k] = a1[i]
		i += 1
		k += 1
	}

	for j < n {
		res[k] = a2[j]
		j += 1
		k += 1
	}

	for i = 0; i < k; i++ {
		a1[i] = res[i]
	}
}