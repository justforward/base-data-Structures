package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		// 检查前k个被保留的元素是否相等（只需要检查num[i-k]即可）
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
