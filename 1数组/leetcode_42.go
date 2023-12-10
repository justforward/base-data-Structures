package main

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	res := 0
	maxIndx := 0
	for i := 1; i < len(height); i++ {
		if height[i] > height[maxIndx] {
			maxIndx = i
		}
	}
	left := height[0]
	for i := 1; i < maxIndx; i++ {
		if height[i] < left { // 可以包住边界
			res += left - height[i]
		} else { // 不能包住边界 直接前移动
			left = height[i]
		}
	}
	right := height[len(height)-1]
	for j := len(height) - 2; j > maxIndx; j-- {
		if height[j] < right {
			res += right - height[j]
		} else {
			right = height[j]
		}
	}
	return res
}
