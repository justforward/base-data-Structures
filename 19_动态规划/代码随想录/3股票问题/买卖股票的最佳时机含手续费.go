package main

func maxProfit6(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		// 不持有的情况下，卖出需要加入手续费
		dp[i][1] = max6(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
		dp[i][0] = max6(dp[i-1][0], dp[i-1][1]-prices[i])
	}
	return dp[n-1][1]
}

func max6(a, b int) int {
	if a > b {
		return a
	}
	return b
}
