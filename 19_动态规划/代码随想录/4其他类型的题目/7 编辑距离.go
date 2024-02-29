package main

/*
  392:判断子序列:判断s是否为t的子序列

这个题目算是编辑距离的入门题目，本题只考虑计算删除的情况，不考虑增加和替换的情况
1、dp[i][j] 表示以下标i-1为结尾的字符串s，和以下标j-1为结尾的字符串t，相同子序列的长度为dp[i][j]
2、递推公式：
   if s[i-1]==t[j-1]
    t中找到了一个字符在s中也出现了
   if s[i-1]!=t[j-1]
    相当于t要删除元素，继续匹配

   if(s[i-1]==t[j-1]),那么dp[i-1][j-1]+1;因为找到了一个相同的字符，相同子序列长度
   自然要在dp[i-1][j-1]+1
   if(s[i-1]!=t[j-1]),此时相当于t要删除元素，t如果把当前元素t[j-1]删除，那么dp[i][j]的数值
	就是看s[i-1]与t[j-2]的比较结果了，即：dp[i][j]=dp[i][j-1]
    其实这里 大家可以发现和 1143.最长公共子序列 (opens new window)的递推公式基本那就是一样的，区别就是 本题 如果删元素一定是字符串t，而 1143.最长公共子序列 是两个字符串都可以删元素。
3、初始化
从递推公式可以看出dp[i][j]都是依赖于dp[i - 1][j - 1] 和 dp[i][j - 1]，所以dp[0][0]和dp[i][0]是一定要初始化的。

这里大家已经可以发现，在定义dp[i][j]含义的时候为什么要表示以下标i-1为结尾的字符串s，和以下标j-1为结尾的字符串t，相同子序列的长度为dp[i][j]。
*/

func isSubsequence(s string, t string) bool {
	len1 := len(t)
	len2 := len(s)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	if dp[len1][len2] == len2 {
		return true
	} else {
		return false
	}

}

/*
 115 不同的子序列：给定一个字符串s和一个字符串t，计算在s中t出现的个数
    这个操作也是只有删除操作

1、dp[i][j] 以i-1为结尾的s子序列中出现j-1为结尾的t的个数
2、递推公式：
	s[i-1]==t[j-1]
       dp[i][j]=dp[i-1][j-1]+dp[i-1][j]
     分为两种情况：
     一部分是用s[i-1]来匹配，那么个数还是为上一个状态的个数 dp[i-1][j-1]
     一部分是不用s[i-1]进行匹配，个数为dp[i-1][j]
     为什么还要考虑 不用s[i - 1]来匹配，都相同了指定要匹配啊？
     例如： s：bagg 和 t：bag ，s[3] 和 t[2]是相同的，但是字符串s也可以不用s[3]来匹配，即用s[0]s[1]s[2]组成的bag。

    s[i-1]!=t[j-1]
       dp[i][j]=dp[i-1][j]
    不相等的情况下，模拟删除s的操作，因为需要在s中找到t出现的个数，所以只需要删除s即可
3、初始化
   dp[i][0]表示什么呢？
   dp[i][0] 表示：以i-1为结尾的s可以随便删除元素，出现空字符串的个数。
   那么dp[i][0]一定都是1，因为也就是把以i-1为结尾的s，删除所有元素，出现空字符串的个数就是1。

	再来看dp[0][j]，dp[0][j]：空字符串s可以随便删除元素，出现以j-1为结尾的字符串t的个数。
    那么dp[0][j]一定都是0，s如论如何也变成不了t。

    dp[0][0]应该是1，空字符串s，可以删除0个元素，变成空字符串t。
*/

func numDistinct(s string, t string) int {
	sLen := len(s)
	tLen := len(t)

	dp := make([][]int, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]int, tLen+1)
	}
	for i := 0; i < sLen; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= tLen; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[sLen][tLen]
}

/*
 1、dp[i][j] 以i-1为结尾的字符串word1和以j-1为结尾的字符串word2，想要达到相等，所需要删除元素的最少次数
2、递推公式
word1[i-1]==word2[j-1]
dp[i][j]=dp[i-1][j-1]
word1[i-1]!=word2[j-1]
2.1) 删除word1[i-1] 最少操作次数为dp[i-1][j]+1
2.2) 删除word2[j-1] 最少操作次数为dp[i][j-1]+1
2.3) 同时删除word1[i-1]和word2[j-1] 操作的最少次数为dp[i-1][j-1]+2

3、dp 如何初始化
dp[i][0]：word2为空字符串，以i-1为结尾的字符串word1要删除多少个元素，才能和word2相同呢，很明显dp[i][0] = i。
同理dp[0][j]=j

*/

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	//初始化
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(word2)+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+2)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
