//https://leetcode.com/problems/fruit-into-baskets/description/
func totalFruit(fruits []int) int {
	fruitMap := make(map[int]int)

	//sliding window
	start, end, max := 0, 0, 0

	for ; end < len(fruits); end++ {
		fruitMap[fruits[end]]++
		for len(fruitMap) > 2 {
			fruitMap[fruits[start]]--
			if fruitMap[fruits[start]] == 0 {
				delete(fruitMap, fruits[start])
			}
			start++
		}
		max = findMax(max, end-start+1)
	}
	return max
}

func findMax(x, y int) int {
	if x >= y {
		return x
	}
	return y
}