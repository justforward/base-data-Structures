package main

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	num2 := []int{2, 5, 6}
	n := 3
	merge(num1, m, num2, n)
	for _, num := range num1 {
		println(num)
	}
}

// 使用从后往前插的思想
func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	index := len(nums1) - 1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[index] = nums1[index1]
			index1--
		} else {
			nums1[index] = nums2[index2]
			index2--
		}
		index--
	}
	for index >= 0 && index1 >= 0 {
		nums1[index] = nums1[index1]
		index--
		index1--
	}
	for index >= 0 && index2 >= 0 {
		nums1[index] = nums2[index2]
		index--
		index2--
	}

}
