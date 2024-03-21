package main

func numIslands(grid [][]byte) int {
	res := 0
	length := len(grid)
	width := len(grid[0])

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j, length, width)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, i, j, length, width int) {
	if i < 0 || j < 0 || i >= length || j >= width {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i+1, j, length, width)
	dfs(grid, i-1, j, length, width)
	dfs(grid, i, j+1, length, width)
	dfs(grid, i, j-1, length, width)

}
