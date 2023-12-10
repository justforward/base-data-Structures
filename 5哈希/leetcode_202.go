package main

// 出现了循环 变换到两个相似的数 就要停下
func isHappy(n int) bool {
	tmpMap := make(map[int]bool, 0)
	for !tmpMap[n] {
		tmpMap[n] = true
		n = sum(n)
		if n == 1 {
			return true
		}
	}
	return false
}

func sum(n int) int {
	res := 0
	for n > 0 {
		res += (n % 10) * (n % 10)
		n = n / 10
	}
	return res
}

func main() {
	println(isHappy(3))

}
