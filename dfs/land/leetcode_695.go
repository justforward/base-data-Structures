package main

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res = max(res, area(grid, i, j))

			}
		}
	}
	return res
}

func area(grid [][]int, i, j int) int {
	if !inArea(grid, i, j) {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 0
	// 所有的岛屿都加起来
	return 1 + area(grid, i+1, j) +
		area(grid, i, j+1) +
		area(grid, i-1, j) +
		area(grid, i, j-1)

}

func inArea(grid [][]int, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
}

func max(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}

func main() {

}
