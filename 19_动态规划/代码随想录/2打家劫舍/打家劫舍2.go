package main

// 首位相连，成环，拆分环，要么只选择从头开始 要么不选择头
func rob_2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res1 := robWithoutCircle(nums[:len(nums)-1])
	res2 := robWithoutCircle(nums[1:len(nums)])
	return max_1(res1, res2)
}

func robWithoutCircle(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max_1(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}

func max_1(i, j int) int {
	if i < j {
		return j
	}
	return i
}
