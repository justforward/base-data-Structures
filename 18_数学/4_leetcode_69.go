package main

/*
 二分查找得到
*/

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	left := 1
	right := x / 2
	for left <= right {
		mid := left + (right-left)/2 // 一定要加上这个
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1
		} else {
			return mid
		}
	}

	//退出的时候right left 返回小于的那个，所以返回right或者left-1
	// 比如当x=8的时候，退出的时候，right=2 left=3
	return right
}

func main() {
	println(mySqrt(8))
}
