package main

/*
1、dp[i] 最长递增子序列的长度，以i为结尾的最长递增子序列的长度、
2、状态转移方程
位置为i的最长子序列=j从0到i-1各个位置的最长升序子序列+1的最大值
if (nums[i] > nums[j]) dp[i] = max(dp[i], dp[j] + 1);
3、dp[i]的初始化
每一个i，对应的dp[i]（即最长递增子序列）起始大小至少都是1.
4、确定遍历顺序
dp[i] 是有0到i-1各个位置的最长递增子序列 推导而来，那么遍历i一定是从前向后遍历。
*/

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
