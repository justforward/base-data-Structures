package main

// 是否为回文数 反转一半，
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10 // 为了得到每一位 然后进行进位操作
		x /= 10                                   //一直在减少最后的一位
	}
	return x == revertedNumber || x == revertedNumber/10
}

func main() {
	x := 123321
	isPalindrome(x)
}
