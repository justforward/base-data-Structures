package main

var res [][]int
var path []int
var used []bool // 记录此时path里面都有哪些元素被使用了，一个排列里面只能使用一次

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	used = make([]bool, len(nums))
	backed(nums)
	return res
}

// 因为排列问题，每次都要从头开始搜索，例如元素1在[1,2]中已经使用过了，但是在[2,1]中还要再使用一次1。
func backed(nums []int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backed(nums)
		used[i] = false
		path = path[:len(path)-1]
	}
}
