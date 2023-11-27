//https://leetcode.com/problems/sort-colors/description/
func sortColors(a []int) {
	start, end, i := 0, len(a)-1, 0

	for i <= end {
		if a[i] == 0 { //always swap with start
			a[i], a[start] = a[start], a[i]
			start++
		}

		if a[i] == 2 { //always swap with end
			a[i], a[end] = a[end], a[i]
			end--
			i--
		}

	}
}