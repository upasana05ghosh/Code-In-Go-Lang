//https://leetcode.com/problems/summary-ranges/description/
func summaryRanges(a []int) []string {
	var res []string
	n := len(a)

	for i := 0; i < n; i++ {
		start := a[i]
		for i+1 < n && a[i]+1 == a[i+1] {
			i++
		}
		if a[i] == start {
			res = append(res, fmt.Sprintf("%d", start))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", start, a[i]))
		}
	}
	return res
}