package leetcode

// DP
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	return maxProfit
}
