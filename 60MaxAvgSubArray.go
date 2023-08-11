//https://leetcode.com/problems/maximum-average-subarray-i/description
func findMaxAverage(a []int, k int) float64 {
	end, sum := 0, 0
	for end < k && end < len(a) {
		sum += a[end]
		end++
	}

	res := sum

	start := 0
	for end < len(a) {
		sum -= a[start]
		start++
		sum += a[end]
		end++
		if res < sum {
			res = sum
		}
	}

	return float64(res) / float64(k)
}