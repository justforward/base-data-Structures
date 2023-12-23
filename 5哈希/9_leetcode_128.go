package main

// n-1 不在哈希表中，才说明这个是连续序列的起点
func longestConsecutive(nums []int) (ans int) {
	set := map[int]bool{}
	for _, n := range nums {
		set[n] = true
	}
	for n := range set {
		if set[n-1] {
			continue
		} // n-1在哈希表中不存在，才说明是连续序列的起点
		cnt := 1 // 以n做起点的连续序列的数量
		for set[n+1] {
			n++
			cnt++
		}
		ans = max(ans, cnt)
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
