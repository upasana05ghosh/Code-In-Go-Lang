//https://leetcode.com/problems/non-overlapping-intervals/description/

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(a, b int) bool { return intervals[a][0] > intervals[b][0] })

	count, pos := 0, intervals[0][0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] > pos { //overlap
			count++
		} else {
			pos = intervals[i][0]
		}
	}

	return count
}