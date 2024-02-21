package main

/*
dp[i][0] 持有股票的状态
dp[i][1] 保持卖出股票的状态
dp[i][2] 卖出股票的状态

dp[i][1]和dp[i][2]在之前的题目中都是一个状态，即不持有股票的状态

dp[i][3] 冷冻期，但是冷冻期状态不可持续，只有一天。

2、确定递推公式
达到买入股票状态（状态一）即：dp[i][0]，有两个具体操作：

操作一：前一天就是持有股票状态（状态一），dp[i][0] = dp[i - 1][0]
操作二：今天买入了，有两种情况
前一天是冷冻期（状态四），dp[i - 1][3] - prices[i]
前一天是保持卖出股票的状态（状态二），dp[i - 1][1] - prices[i]
那么dp[i][0] = max(dp[i - 1][0], dp[i - 1][3] - prices[i], dp[i - 1][1] - prices[i]);

达到保持卖出股票状态（状态二）即：dp[i][1]，有两个具体操作：

操作一：前一天就是状态二
操作二：前一天是冷冻期（状态四）
dp[i][1] = max(dp[i - 1][1], dp[i - 1][3]);

达到今天就卖出股票状态（状态三），即：dp[i][2] ，只有一个操作：

昨天一定是持有股票状态（状态一），今天卖出

即：dp[i][2] = dp[i - 1][0] + prices[i];

达到冷冻期状态（状态四），即：dp[i][3]，只有一个操作：

昨天卖出了股票（状态三）

dp[i][3] = dp[i - 1][2];

*/

func maxProfit5(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp := make([][]int, n)
	status := make([]int, n*4)
	for i := range dp {
		dp[i] = status[:4]
		status = status[4:]
	}
	dp[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max5(dp[i-1][0], max5(dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
		dp[i][1] = max5(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max5(dp[n-1][1], max5(dp[n-1][2], dp[n-1][3]))
}

func max5(a, b int) int {
	if a > b {
		return a
	}
	return b
}
