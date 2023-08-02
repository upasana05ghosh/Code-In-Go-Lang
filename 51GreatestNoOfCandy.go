//https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description

func kidsWithCandies(c []int, e int) []bool {
	max := 0

	for _, v := range c {
		if max < v {
			max = v
		}
	}
	result := make([]bool, len(c))

	for i, v := range c {
		if v+e >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}