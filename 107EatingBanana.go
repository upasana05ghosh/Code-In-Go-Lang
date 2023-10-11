//https://leetcode.com/problems/koko-eating-bananas/description/
func minEatingSpeed(piles []int, h int) int {
	//binary search
	l := 1
	r := math.MinInt32

	for _, v := range piles {
		r = max(r, v)
	}

	for l <= r {
		k := (l + r) / 2
		if eatAll(piles, k, h) {
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return l
}

func eatAll(piles []int, k int, h int) bool {
	count := 0

	for _, v := range piles {
		count += v / k
		if v%k != 0 {
			count++
		}
	}

	return count <= h
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}