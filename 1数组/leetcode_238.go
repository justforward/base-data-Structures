package main

// 原数组：       [1       2       3       4]
//左部分的乘积：   1       1      1*2    1*2*3
//右部分的乘积： 2*3*4    3*4      4      1
//结果：        1*2*3*4  1*3*4   1*2*4  1*2*3*1

// 从左到右遍历一边，把每个数左边的乘积求出来，再从右到左遍历一边，在原来的结果基础上每个数右边的乘积
func productExceptSelf(nums []int) []int {
	size := len(nums)
	ans := make([]int, size)
	left_ans := make([]int, size)
	right_ans := make([]int, size)
	left_ans[0] = 1
	right_ans[size-1] = 1
	// 这个是把最后一个元素空出来 不取到最后一个元素
	for i := 1; i < len(nums); i++ {
		left_ans[i] = left_ans[i-1] * nums[i-1]
	}

	// 这个是把前一个元素空出来 不取到第一个元素
	for i := len(nums) - 2; i >= 0; i-- {
		right_ans[i] = right_ans[i+1] * nums[i+1]
	}
	for j := 0; j < len(nums); j++ {
		ans[j] = left_ans[j] * right_ans[j]
	}

	return ans
}

func main() {

}
