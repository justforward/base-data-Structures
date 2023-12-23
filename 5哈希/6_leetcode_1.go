package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for index, value := range nums {
		if _, ok := m[target-value]; ok {
			res = append(res, index, m[target-value])
			return res
		}
		m[value] = index
	}
	return res
}
