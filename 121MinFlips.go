//https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/description
func minFlips(a int, b int, c int) int {
	res := 0

	for a > 0 || b > 0 || c > 0 {
		if c&1 == 0 {
			res += a&1 + b&1
		} else if a&1 == 0 && b&1 == 0 {
			res++
		}

		//right shift
		a = a >> 1
		b = b >> 1
		c = c >> 1
	}

	return res
}