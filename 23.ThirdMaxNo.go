//https://leetcode.com/problems/third-maximum-number/description/
func thirdMax(nums []int) int {
	m1, m2, m3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {
		if v == m1 || v == m2 || v == m3 {
			continue
		}

		if m3 == math.MinInt64 || v > m3 {
			m1 = m2
			m2 = m3
			m3 = v
		} else if m2 == math.MinInt64 || v > m2 {
			m1 = m2
			m2 = v
		} else if m1 == math.MinInt64 || v > m1 {
			m1 = v
		}
	}

	if m1 != math.MinInt64 {
		return m1
	}
	return m3
}