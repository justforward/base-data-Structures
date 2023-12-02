package main

func minSubArrayLen(target int, nums []int) int {
	res, sum := len(nums)+1, 0
	left, right := 0, 0
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			sublen := (right - left) + 1
			if sublen < res {
				res = sublen
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if res == len(nums)+1 {
		return 0
	} else {
		return res
	}
}
