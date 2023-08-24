//https://leetcode.com/problems/increasing-triplet-subsequence/description/
func increasingTriplet(a []int) bool {
	min := math.MaxInt32
	max := math.MaxInt32

	for _, v := range a {
		if v <= min {
			min = v
		} else if v <= max {
			max = v
		} else {
			return true
		}
	}

	return false
}