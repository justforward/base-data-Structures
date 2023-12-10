package main

// 1、使用一个额外的数组，将原数组结果放入到这个额外的数组中，然后再copy到原来的数组中 空间复杂度为O（n）
// 2、空间复杂度为O（1）
// 2.1）首先对整个数组实行翻转，这样子原数组中需要翻转的子数组，就会跑到数组最前面。
// 2.2）这时候，从 kkk 处分隔数组，左右两数组，各自进行翻转即可。
func rotate_1(nums []int, k int) {
	tmp := make([]int, len(nums))
	for index, value := range nums {
		tmp[(index+k)%len(nums)] = value
	}

	// 把tmp复制给nums
	copy(nums, tmp)
}

func rotate_2(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(n []int) {
	size := len(n)
	for i, j := 0, size; i < size/2; i++ {
		n[i], n[j-1-i] = n[j-1-i], n[i]
	}
}
func main() {

}
