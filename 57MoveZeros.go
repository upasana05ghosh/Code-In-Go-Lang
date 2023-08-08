//https://leetcode.com/problems/move-zeroes/description/
func moveZeroes(a []int) {
	slow := 0
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			a[slow] = a[i]
			slow++
		}
	}

	for slow < len(a) {
		a[slow] = 0
		slow++
	}
}