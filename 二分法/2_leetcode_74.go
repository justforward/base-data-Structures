package main

func searchMatrix(matrix [][]int, target int) bool {
	// 从左上方开始查找
	first, second := len(matrix)-1, 0
	for first >= 0 && second < len(matrix[0]) {
		if target > matrix[first][second] {
			second++
		} else if target < matrix[first][second] {
			first--
		} else {
			return true
		}
	}
	return false
}
