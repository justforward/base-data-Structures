package main

// 数组中的第k个最大数
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, index int) int {
	if l == r {
		return nums[l]
	}

	pivot := nums[l]
	i := l + 1
	j := r

	for {
		for i < r && nums[i] <= pivot {
			i++
		}
		for j > l && nums[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[l], nums[j] = nums[j], nums[l]

	if index == j {
		return nums[j]
	} else if index < j {
		return quickSelect(nums, l, j-1, index)
	} else {
		return quickSelect(nums, j+1, r, index)
	}
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	println(findKthLargest(nums, 4))
}
