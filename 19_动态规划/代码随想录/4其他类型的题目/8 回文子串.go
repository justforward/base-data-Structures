package main

/*

计算这个字符串中有多少个回文子串
1、dp[i][j] 表示区间范围[i,j]的子串是否为回文子串，如果是dp[i][j]为true，否则为false
2、递推公式：
if s[i]!=s[j]
dp[i][j]=false
if s[i]==s[j]
2.1) aa j-i<=1 dp[i][j]=true
2.2) a i j 指向同一个字符 ｜ cabac 需要查看dp[i+1][j-1]是否相同
3、初始化都为false
4、遍历顺序，看到dp[i][j] 是通过dp[i+1][j-1]推到得到
 dp[i + 1][j - 1] 在 dp[i][j]的左下角

dp[i][j-1]     dp[i][j]
dp[i+1][j-1]   dp[i+1][j]

所以一定要从下到上，从左到右遍历，这样保证dp[i + 1][j - 1]都是经过计算的。

*/

func countSubstrings(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}
	return res
}
