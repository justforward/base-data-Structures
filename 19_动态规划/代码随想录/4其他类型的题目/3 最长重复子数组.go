package main

/*
1、dp[i][j]:表示以i-1为结尾的num1，以j-1为结尾的num2 最长子数组长度
2、递推公式：
	if nums1[i-1]==nums2[j-1]{
	    dp[i][j]=dp[i-1][j-1]+1
     }
    递推公式里面，为什么不是dp[i][j-1]或者dp[i-1][j]? 一定是dp[i-1][j-1]?
3、初始化
	根据dp[i][j]的定义，dp[i][0] 和dp[0][j]其实都是没有意义的！
	为什么没有意义？因为dp[i][0]表示以i-1，0-1 为结尾的数组，超出了数组下表的限制
	但dp[i][0] 和dp[0][j]要初始值，因为 为了方便递归公式dp[i][j] = dp[i - 1][j - 1] + 1;
	所以dp[i][0] 和dp[0][j]初始化为0。
	举个例子A[0]如果和B[0]相同的话，dp[1][1] = dp[0][0] + 1，只有dp[0][0]初始为0，正好符合递推公式逐步累加起来。
4、确定遍历顺序
外层for循环遍历A，里层for循环遍历B。
同时题目要求长度最长的子数组的长度。所以在遍历的时候顺便把dp[i][j]的最大值记录下来。
for i:=1;i<=len(nums);i++
	for j:=1;j<=len(nums);j++
为什么需要取到最后一个？主要是以i-1为结尾的num1，
*/

func findLength(A []int, B []int) int {
	m, n := len(A), len(B)
	res := 0
	dp := make([][]int, m+1) //i-1为结尾 所以需要
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

/*
	就定义dp[i][j]为 以下标i为结尾的A，和以下标j 为结尾的B，最长重复子数组长度
1、初始化的时候麻烦：
如果定义 dp[i][j]为 以下标i为结尾的A，和以下标j 为结尾的B，
那么 第一行和第一列毕竟要进行初始化，
如果nums1[i] 与 nums2[0] 相同的话，对应的 dp[i][0]就要初始为1，
因为此时最长重复子数组为1。 nums2[j] 与 nums1[0]相同的话，同理。
2、遍历的时候，
	for (int i = 0; i < nums1.size(); i++) {
            for (int j = 0; j < nums2.size(); j++) {
*/
