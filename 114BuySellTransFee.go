//https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
func maxProfit(p []int, fee int) int {
	sell := 0
	buy := -p[0]

	for i := 1; i < len(p); i++ {
		prevBuy := buy
		buy = max(buy, sell-p[i])
		sell = max(sell, prevBuy+p[i]-fee) //selling
	}

	return sell
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}