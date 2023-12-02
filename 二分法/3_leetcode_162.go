package main

// 因为给到的题目中说明了，-1 和num[len(nums)] 等于负无穷说明 两端一定会往下降落，不用担心二分找到递增区间
// 为什么选用左闭右开的？根题目使用的条件有关系。
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1 // 虽然是左闭右开的区间，但是因为要取到mid+1 所以right不能取到len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] { // 如果取到left==right的时候 就没有办法取到这个位置的数据
			right = mid // 右开区间，右边没有取到，下一次就需要取到
		} else {
			left = mid + 1 //左闭区间，左边已经取到，下一次就不需要取到
		}
	}
	return left // 返回right也是一样的，停止条件left==right
}
