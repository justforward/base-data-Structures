package main

/*
1) dp[i] 以nums[i]为结尾的 最大连续子序列和为dp[i]
2) 递推公式：只能由两个方向推到出来
   2.1）dp[i-1]+nums[i],从nums[i] 加入当前连续子序列和
   2.2）nums[i],从头开始计算连续子序列和
	这两个方向求max
3）初始化
dp[0]就是递推公式的基础
dp[0]=nums[0]
*/

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxSub(dp[i-1]+nums[i], nums[i])
		ans = maxSub(ans, dp[i])
	}
	return ans
}

func maxSub(i, j int) int {
	if i < j {
		return j
	}
	return i
}
