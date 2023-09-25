//https://leetcode.com/problems/single-number/description/
func singleNumber(a []int) int {
	xor := 0
	for _, v := range a {
		xor ^= v
	}

	return xor
}