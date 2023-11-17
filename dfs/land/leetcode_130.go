package main

// 从最边缘开始找到o 并找到与之相连的o区域，标记下来，然后把剩下的o全部替换为x即可
// 将被标记过的边缘o=A

func solve(board [][]byte) {
	// 从最边缘开始找到所有于边界
	for i := 0; i < len(board); i++ {
		dfs(board, i, 0)
		dfs(board, i, len(board[0])-1)
	}

	for j := 1; j < len(board[0])-1; j++ {
		dfs(board, 0, j)
		dfs(board, len(board)-1, j)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y+1)
	dfs(board, x, y-1)
}

func main() {

}
