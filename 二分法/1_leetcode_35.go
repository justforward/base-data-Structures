package main

// 第一种左闭右闭的区间
// 最终停下里的位置是 【right，left】 那么left 和right+1就是要插入的元素点
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	println(left, right)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}

func main() {
	nums := []int{3, 1}
	println(searchInsert(nums, 1))
}
