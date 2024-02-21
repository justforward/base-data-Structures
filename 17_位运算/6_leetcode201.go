package main

// 给定两个整数，我们要找到他们对应的二进制字符串的公共前缀
func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right &= (right - 1)
	}
	return right
}
