package main

// 快指针：寻找新数组的元素，新数组就是不含目标元素的数组
// 慢指针：指向更新 新数组下标的位置
func removeElement(nums []int, val int) int {
	fast := 0
	slow := 0
	for fast < len(nums) {
		if val != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow

}
