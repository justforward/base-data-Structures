package main

func search(nums []int, target int) int {
	//已经保证这个数组里面的值互不相同
	//首先找到旋转点
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= nums[0] {
			//判断右边是否同样情况
			l = mid + 1
		} else if nums[mid] < nums[0] { // l 是一个转折点 l==r的时候 下一个要更新的位置跳出
			r = mid - 1 // r -1
		}
	}
	println("get left: ", l, "right: ", r)
	l = r // r指向转折点的前一个点
	//确定这个数的查询在哪个递增区间里面
	if target >= nums[0] {
		l = 0
	} else {
		l = l + 1
		r = len(nums) - 1
	}
	println("get left: ", l, "right: ", r)

	return searchNum(nums, l, r, target)

}

func searchNum(nums []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	println(search(nums, 1))
}
