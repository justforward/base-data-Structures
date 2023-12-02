package main

// 1、情况一 左值<中值，中值>右值，直接收缩右边界,中值在右区间内？
// 2、情况二，左值>中值，中值<右值，直接
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right - left) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		}
	}
	return nums[left]
}
