package main

import "sort"

var (
	path []int
	res  [][]int
)

func combinationSum(candidates []int, target int) [][]int {
	path = make([]int, 0)
	res = make([][]int, 0)
	sort.Ints(candidates) // 一定要排序，否则会出现错误剪枝的情况
	dfs(candidates, 0, target, 0)
	return res
}

func dfs(candidates []int, index, target int, sum int) {
	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}

	for i := index; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			break
		}
		path = append(path, candidates[i])
		dfs(candidates, i, target, sum+candidates[i])
		path = path[:len(path)-1]
	}
}
