package main

func lengthOfLastWord(s string) int {
	var ans int
	for i := len(s) - 1; i >= 0; i-- {
		// 当前字符是空格，且之前包含非空格
		if s[i] == ' ' && ans > 0 {
			break
		}
		// 空格跳过
		// 非空格+1
		if s[i] != ' ' {
			ans++
		}
	}
	return ans
}
