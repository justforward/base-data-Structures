package main

// 找到最大k个数 使用快排实现

func findKthLargest(nums []int, k int) int {
	res := findIndex(nums, 0, len(nums)-1, len(nums)-k+1)
	return res
}

func findIndex(nums []int, left, right, k int) int {
	if left > right {
		return -1
	}
	index := quickSort(nums, left, right)
	if index+1 > k {
		//查找左边
		return findIndex(nums, left, index-1, k)
	} else if index+1 < k {
		return findIndex(nums, index+1, right, k)
	} else {
		return nums[index]
	}

}

func quickSort(nums []int, left, right int) int {
	key := nums[left]
	for left < right {
		for left < right && nums[right] >= key {
			right--
		}
		swap(nums, left, right)
		for left < right && nums[left] <= key {
			left++
		}
		swap(nums, left, right)
	}
	return left
}

func swap(num []int, left, right int) {
	key := num[left]
	num[left] = num[right]
	num[right] = key
}
