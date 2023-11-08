//https://leetcode.com/problems/jump-game/description/
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	n := len(nums)
	a := make([]int, n)
	a[0] = 1
	i := 0
	var j int
	for i < n {
		if a[i] == 1 {
			j = endReached(i+1, nums[i], a)
			if j == n {
				return true
			}
		} else {
			return false
		}
		i++
	}

	if a[i-1] == 1 {
		return true
	}
	return false
}

func endReached(i, count int, a []int) int {
	for count > 0 && i < len(a) {
		a[i] = 1
		count--
		i++
	}

	return i
}