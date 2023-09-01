//https://leetcode.com/problems/find-the-highest-altitude/description/
func largestAltitude(a []int) int {
	max, b := 0, 0

	for _, v := range a {
		b = b + v
		if b > max {
			max = b
		}
	}

	return max
}