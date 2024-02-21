package main

/*
   1、dp[i] 字符串长度为i的时候，dp[i]=true ,表示可以拆分一个或者多个在字典中出现的单词
   2、

   求组合数就是外层for循环遍历物品，内层for循环遍历背包（不要求有顺序）
   求排列数就是外层for遍历背包，内层for循环遍历物品 （要求有顺序）
*/
// 求装满背包s的前几位字符的方式有几种
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)

	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ { // 需要截取[j:i]
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
