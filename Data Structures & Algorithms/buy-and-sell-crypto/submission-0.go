func maxProfit(prices []int) int {
	maxProfit := 0
	minBuy := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < minBuy {
			minBuy = prices[i]
		}
		profit := prices[i] - minBuy
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
