package main

// O(n)的方式解决
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right, k := 0, len(nums)-1, len(nums)-1

	for left <= right {
		leftNums := nums[left] * nums[left]
		rightNums := nums[right] * nums[right]
		if leftNums < rightNums {
			res[k] = rightNums
			right--
		} else {
			res[k] = leftNums
			left++
		}
		k--
	}
	return res

}
