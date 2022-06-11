package problem121

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	ans := 0
	min := 1 << 30
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			ans = max(ans, prices[i]-min)
		}
	}
	return ans
}
