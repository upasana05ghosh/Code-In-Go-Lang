//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func maxProfit(prices []int) int {
	profit := 0
	if len(prices) == 0 {
		return profit
	}
	min, max := prices[0], prices[0]

	for _, price := range prices {
		if min > price {
			min = price
			//reset max
			max = min
		}
		max = findMax(max, price)
		profit = findMax(profit, max-min)
	}

	return profit
}

func findMax(x, y int) int {
	if x >= y {
		return x
	}
	return y
}