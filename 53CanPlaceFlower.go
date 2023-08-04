//https://leetcode.com/problems/can-place-flowers/description
func canPlaceFlowers(bed []int, n int) bool {
	count := 1 //initializing for 1st element
	result := 0

	for i := 0; i < len(bed); i++ {
		if bed[i] == 0 {
			count++
		} else {
			result += (count - 1) / 2
			count = 0
		}
	}

	if count != 0 {
		result += count / 2
	}

	return result >= n
}
