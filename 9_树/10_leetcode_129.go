package main

func sumNumbers(root *TreeNode) int {
	return dfs4(root, 0)
}

func dfs4(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	} else {
		return dfs4(root.Left, sum) + dfs4(root.Right, sum)
	}
}
