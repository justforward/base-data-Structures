package main

// left right
// 移动的条件：比较低的那块开始移动
func maxArea(height []int) int {
	res, left, right := 0, 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > res {
			res = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
