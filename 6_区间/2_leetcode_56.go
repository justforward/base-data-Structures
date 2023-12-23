package main

import "sort"

func merge(intervals [][]int) (ans [][]int) {
	// 步骤1：按照每个区间的起始值进行升序排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	ans = append(ans, intervals[0]) // 初始化结果数组，将排序后的第一个区间加入其中

	// 步骤3：遍历排序后的区间数组，从第二个区间开始
	for _, e := range intervals[1:] {
		// 步骤4：检查当前区间与结果数组中最后一个区间是否重叠
		if ans[len(ans)-1][1] < e[0] {
			// 如果没有重叠，将当前区间加入结果数组
			ans = append(ans, e)
		} else {
			// 如果有重叠，更新结果数组中最后一个区间的结束值为当前区间的结束值
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], e[1])
		}
	}

	// 步骤5：返回合并后的结果数组
	return ans
}

// max 函数用于比较两个整数，返回较大的一个
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
