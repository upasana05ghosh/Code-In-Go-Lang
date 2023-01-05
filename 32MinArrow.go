//https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/
func findMinArrowShots(points [][]int) int {
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][0] > points[j][0]
	})

	arrCount := 1
	pos := points[0][0]
	i := 1

	for i < len(points) {
		if points[i][1] < pos {
			arrCount += 1
			pos = points[i][0]
		}
		i++
	}
	return arrCount
}