package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for index, n := range nums {
		if value, ok := m[n]; ok {
			if abs(index, value) <= k {
				return true
			}
		}
		m[n] = index
	}
	return false
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	} else {
		return b - a
	}
}
