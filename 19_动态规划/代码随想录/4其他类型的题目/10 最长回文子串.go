package main

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 { // If the input string is less than 2, directly return s
		return s
	}

	start, maxLen := 0, 1 //记录长度即可
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
			// 一定在完成之后判断
			if dp[i][j] && j-i+1 > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}

	return s[start : start+maxLen]
}
