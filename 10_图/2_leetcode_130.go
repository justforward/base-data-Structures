package main

/*
	找中间的联通部分，比如从两边找不用变X的所有O，那么剩下的所有O都是要变X的

	被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。
   任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。
*/

func solve(board [][]byte) {
	// 从最边缘开始找到所有于边界
	for i := 0; i < len(board); i++ {
		dfs_1(board, i, 0)
		dfs_1(board, i, len(board[0])-1)
	}

	for j := 1; j < len(board[0])-1; j++ {
		dfs_1(board, 0, j)
		dfs_1(board, len(board)-1, j)
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

func dfs_1(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs_1(board, x+1, y)
	dfs_1(board, x-1, y)
	dfs_1(board, x, y+1)
	dfs_1(board, x, y-1)
}
