package main

/*

1、dp[i][0]	表示第i天持有股票所得最多现金，dp[i][1]表示第i天不持有股票所得最多现金
    注意这里说的是“持有”，“持有”不代表就是当天“买入”！也有可能是昨天就买入了，今天保持持有的状态
2、确定递推公式
如果第i天持有股票dp[i][0],那么可以有两个状态推到出来:
1）第i-1天就持有股票那么就维持现状，所得现金就是昨天持有股票的所得现金，
2）第i天买入股票，所得现金就是买入今天的股票所得现金：-prices[i]
那么dp[i][0]=max(dp[i-1][0],-prices[i])

如果第i天不持有股票dp[i][1],那么也可以由两个状态推出:
1)第i-1天不持有股票，那就保持现状，dp[i-1][1]
2) 第i天卖出，所得现金就是按照今天价格卖出去的现金：prices[i]+dp[i-1][0]
dp[i][1]=max(dp[i-1][1],prices[i]+dp[i-1][0])

3、dp数组如何初始化
其基础都是要从dp[0][0]和dp[0][1]推导出来。
dp[0][1]=0 第0天不持有股票
dp[0][0] -= prices[0]; 第0天持有股票

4、确定遍历顺序
从递推公式可以看出dp[i]都是由dp[i - 1]推导出来的，那么一定是从前向后遍历。

5、最终结果是：dp[最后一天][1] 因为本题中在最后一天不持有股票状态所得金钱一定比持有股票状态得到的多！
*/

func maxProfit(prices []int) int {
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
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
