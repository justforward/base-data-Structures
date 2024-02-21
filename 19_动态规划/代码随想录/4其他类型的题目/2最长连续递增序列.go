package main

/*
1、 dp[i] 为以i为结尾的最长连续递增序列
2、要求连续，所以只能对比前一个nums[i-1]<nums[i] 的时候 dp[i]=dp[i-1]+1
*/

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return ans
}
