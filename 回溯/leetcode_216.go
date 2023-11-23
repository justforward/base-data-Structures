package main

//
//var (
//	path []int
//	res  [][]int
//)
//
//func combinationSum3(k int, n int) [][]int {
//	path = make([]int, 0)
//	res = make([][]int, 0)
//	dfs(k, n, 0, 1)
//	return res
//}
//
//func dfs(k, targetSum, sum, startIndex int) {
//
//	if len(path) == k {
//		if targetSum == sum {
//			tmp := make([]int, k) // 这里为什么要一定设置大小？ copy的底层实现：从源切片/数组的开头开始，逐个复制元素到目标切片/数组，直到达到两者中长度较短的那一个。
//			copy(tmp, path)
//			res = append(res, tmp)
//			return
//		}
//	}
//
//	for i := startIndex; i <= 9; i++ {
//		if sum+i > targetSum || 9-i+1 < k-len(path) {
//			break
//		}
//		path = append(path, i)
//		dfs(k, targetSum, sum+i, i+1)
//		path = path[:len(path)-1]
//	}
//}
