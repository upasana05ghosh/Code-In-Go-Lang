//https://leetcode.com/problems/maximum-ice-cream-bars/description/
func maxIceCream(costs []int, coins int) int {
	//sort.Ints(costs)

	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})
	count := 0
	for _, v := range costs {
		if coins < v {
			break
		}
		coins -= v
		count++
	}
	return count
}