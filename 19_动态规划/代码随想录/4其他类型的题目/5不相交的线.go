package main

/*
 最长公共子序列的长度
1）dp[i][j] 表示下标从0～i-1 nums1和下标从0～j-1 nums2相同的公共子序列长度
2）递推公式：
	if nums1[i-1]==nums2[j-1]:
		dp[i][j]=dp[i-1][j-1]+1
    else：
        dp[i][j]=max(dp[i-1][j],dp[i][j-1])
3）初始化：
dp[i][0]=0
dp[0][j]=0
*/

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	t1 := len(nums1)
	t2 := len(nums2)
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i <= t1; i++ {
		for j := 1; j <= t2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxCurr(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[t1][t2]
}

func maxCurr(i, j int) int {
	if i < j {
		return j
	}
	return i
}
