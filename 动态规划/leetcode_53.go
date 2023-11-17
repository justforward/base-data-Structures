package main

// dp[i]=以下标为i的连续子数组最大和
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		curr := dp[i-1] + nums[i]
		dp[i] = max(curr, nums[i])
		res = max(dp[i], res)
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {

}
