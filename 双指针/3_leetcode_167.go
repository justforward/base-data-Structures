package main

func twoSum(numbers []int, target int) []int {
	res := []int{-1, -1}
	left, right, sum := 0, len(numbers)-1, 0
	for left < right {
		sum = numbers[left] + numbers[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			res[0] = left + 1
			res[1] = right + 1
			return res
		}
	}
	return res
}
