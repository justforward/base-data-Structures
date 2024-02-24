package main

/*
 为什么这里是0～i-1 跟连续子数组为结尾的不一样。
上一题以i-1和j-1是因为子数组必须要求是连续的，如果不连续，公共子数组直接归零，
下一个子数组不能继承前一个子数组的公共子数组长度。
子序列则不一样，允许中间有间隔，
下一个子序列可以继承前一个子序列的公共子序列长度。
这样说很抽象，我们举个例子。比如说两个数组nums1 = 【1,2,3,4,5】 , nums2 = 【1,2,3,8,5】
。在index=3的时候出现分歧了，
如果是公共子数组，到index=3时，其公共子数组必须要归零，如果不归零，会影响index=4的判断。
而如果是公共子序列，index=3可以保留index=2的最长子序列数，继而在index=4时继续递增，
主要问题就在于，如果你想要这个又要的递推公式dp【i】【j】 = dp【i-1】【j-1】 + 1
就必须这样定义，你当然也可以用上一题的方式定义这一题，或者用这一题的方式定义上一题，但你就需要一个非常麻烦的递推公式

1、dp[i][j]：长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的最长公共子序列为dp[i][j]
 2、递推公式
	if num1[i-1]==num2[j-1]
       dp[i][j]=dp[i-1][j-1]+1
    else
       dp[i][j]=max(dp[i][j-1],dp[i-1][j])

  举例：
	abc
	abe   c和e不等，但是可以考虑abc和ab的最长公共子序列，或者考虑ab和abe的最长公共子序列
          但是在上一个题目中，只能考虑连续的最长公共子数组，所以都需要往后退一个，否则就被清空

*/

func longestCommonSubsequence(text1 string, text2 string) int {
	t1 := len(text1)
	t2 := len(text2)
	dp := make([][]int, t1+1)
	for i := range dp {
		dp[i] = make([]int, t2+1)
	}

	for i := 1; i <= t1; i++ {
		for j := 1; j <= t2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max1(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[t1][t2]
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
