package main

// dp[i] 偷当前房子获取的最大额度 这个为什么需要dp
func rob_1(nums []int) int {
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
	}
	return i
}
