package main

import (
	"fmt"
	"strconv"
)

// 分组循环
// 适用场景：数组被分成诺干段，每一段的逻辑处理是一样的
func summaryRanges(nums []int) []string {
	ans := []string{}
	i, n := 0, len(nums)
	for i < n {
		start := i
		for i < n-1 && nums[i]+1 == nums[i+1] {
			i++
		}
		s := strconv.Itoa(nums[start])
		if start < i {
			s += "->" + strconv.Itoa(nums[i])
		}
		ans = append(ans, s)
		i++ // 这里一定是要++的 否则一直在当前坐标中
	}
	return ans
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}
