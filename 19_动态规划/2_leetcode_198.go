package main

// dp[i]到当前房间偷盗的最高金额 1 2 3 4 dp[i]=max(dp[i-2]+num[i],dp[i-1])
func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}

func max(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}
