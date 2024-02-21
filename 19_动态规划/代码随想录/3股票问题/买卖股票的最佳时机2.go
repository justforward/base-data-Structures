package main

/*
相对于只有一次买卖，这正是因为本题的股票可以买卖多次！所以买入股票的时候，可能会有之前买卖的利润即：dp[i - 1][1]，所以dp[i - 1][1] - prices[i]。
*/
func maxProfit2(prices []int) int {

	length := len(prices)
	if length == 0 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < length; i++ {
		// 之前的是 max(dp[i-1][0], -prices[i]) 这个是需要继承原来的，本题的股票可以买卖多次！ 所以买入股票的时候，可能会有之前买卖的利润。
		dp[i][0] = max2(dp[i-1][0], dp[i-1][1]-prices[i]) // // 注意这里是和121. 买卖股票的最佳时机唯一不同的地方
		dp[i][1] = max2(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

func max2(i, j int) int {
	if i < j {
		return j
	}
	return i
}
