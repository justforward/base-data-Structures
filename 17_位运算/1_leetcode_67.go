package main

import "strconv"

// 采取末尾对齐 进行累加 然后拼接结果
func addBinary(a string, b string) string {
	ans := ""
	carray := 0
	lenA := len(a)
	lenB := len(b)
	n := max(lenA, lenB)
	for i := 0; i < n; i++ {
		if i < lenA {
			carray += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carray += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carray%2) + ans
		carray /= 2
	}
	if carray > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}
