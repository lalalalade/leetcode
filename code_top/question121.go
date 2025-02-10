package code_top

// maxProfit 买卖股票的最佳时机
func maxProfit(prices []int) int {
	res := 0
	minPrice := prices[0]
	for _, p := range prices {
		res = max(res, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return res
}
