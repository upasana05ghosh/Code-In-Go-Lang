//https://leetcode.com/problems/container-with-most-water/description/
func maxArea(a []int) int {
	//wider distance meaning more water
	i, j, max := 0, len(a)-1, 0

	for i < j {
		h := findMin(a[i], a[j])
		max = findMax(max, h*(j-i))

		//Move smaller one
		if a[i] < a[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}