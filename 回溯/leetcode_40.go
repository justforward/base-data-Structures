package main

/*import "sort"

// 每个元素在同一个组合中只能使用一次，但是这个元素可能有重复的值
// 比如存在元素 1 1 1 1 等到的结果为2 那么组合【1 1】 【1 1】

// 解题的核心：相同层的相同元素不能使用 去掉中间两个1 1的情况

// 回看一下题目，元素在同一个组合内是可以重复的，怎么重复都没事，但两个组合不能相同。
// 强调一下，树层去重的话，需要对数组排序！

//如果candidates[i] == candidates[i - 1] 并且 used[i - 1] == false，
//就说明：前一个树枝，使用了candidates[i - 1]，也就是说同一树层使用过candidates[i - 1]。

var (
	used []bool
	path []int
	res  [][]int
)

func combinationSum2(candidates []int, target int) [][]int {
	used = make([]bool, len(candidates))
	path = make([]int, 0)
	res = make([][]int, 0)
	sort.Ints(candidates)
	dfs(candidates, 0, target, 0)
	return res
}

func dfs(candidates []int, index int, target int, sum int) {
	if target == sum {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := index; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			break
		}
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		if i > 0 && candidates[i-1] == candidates[i] && used[i-1] == false {
			continue
		}

		path = append(path, candidates[i])
		// 这种取值的时候 一定要注意初始化数组的时候，需要加上长度
		used[i] = true
		dfs(candidates, i+1, target, sum+candidates[i])
		used[i] = false
		path = path[:len(path)-1]
	}
}*/
