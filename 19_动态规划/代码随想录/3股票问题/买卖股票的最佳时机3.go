package main

/*
关键在于至多买卖两次，这意味着可以买卖一次，可以买卖两次，也可以不买卖。
1、确定dp数组以及对应的下标含义
	一天一共就有五个状态，

	0）没有操作 （其实我们也可以不设置这个状态）
	1）第一次持有股票
 	2）第一次不持有股票
	3）第二次持有股票
	4）第二次不持有股票

dp[i][j]中表示第i天，j表示0-4五个状态，dp[i][j]表示第i天状态j所剩的最大现金。
需要注意：dp[i][1]，表示的是第i天，买入股票的状态，并不是说一定要第i天买入股票，这是很多同学容易陷入的误区。
例如 dp[i][1] ，并不是说 第i天一定买入股票，有可能 第 i-1天 就买入了，
那么 dp[i][1] 延续买入股票的这个状态。

2、确定递推公式
达到dp[i][1]状态，有两个具体操作：
1）第i天买入了股票，那么dp[i][1]=dp[i-1][0]-prices[i]
2) 第i天没有操作，沿用前一天买入的状态，dp[i][1]=dp[i-1][1]
两个取max
dp[i][2]，也存在两个操作：
1）在第i天进行交易，卖出股票。dp[i][2]=dp[i-1][1]+prices[i]
2) 在第i天没有操作，而是沿用前一天卖出的状态：dp[i][2]=dp[i-1][2]

dp[i][3] = max(dp[i - 1][3], dp[i - 1][2] - prices[i]);
dp[i][4] = max(dp[i - 1][4], dp[i - 1][3] + prices[i]);

3、初始化
dp[0][1]=-prices[0]
dp[0][3]=-prices[0]
*/

func maxProfit3(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(prices)-1][4]
}
