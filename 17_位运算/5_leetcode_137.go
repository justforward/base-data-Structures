package main

// 统计每个比特位上1的个数，或者用位运算实现mod3 加法
func singleNumbers(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		cnt1 := int32(0)
		for _, x := range nums {
			cnt1 += int32(x) >> i & 1 // 得到每个数字上1的个数
		}
		ans = ans | cnt1%3<<i // 对3取模的结果移动i位，就是把这个结果变到原来的位置，然后｜操作是将结果相加
	}
	return int(ans)
}
