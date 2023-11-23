package main

// 1、哈希统计各数字的数量，找到众数，时间复杂度和空间复杂度均为O(n)
// 2、摩尔投票法：核心理念为票数正负相抵，时间和空间复杂度为O（n）和O（1）
func majorityElement_1(nums []int) int {
	tmpMap := make(map[int]int)
	size := len(nums) / 2

	for _, n := range nums {
		if tmpMap[n] == size-1 {
			return n
		} else {
			tmpMap[n]++
		}
	}
	return 0
}

// 投票算法
// 当count==0 说明当前major空缺
// 当count！=0 说明当前major的票数没有被完全抵消，所以让count--
func majorityElement_2(nums []int) int {
	major := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}

		if major == num {
			count++
		} else {
			count--
		}
	}
	return major
}

func main() {

}
