package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 本题目可以使用前序，也可以使用后序变细，使用前序列遍历求的是深度，后序求的是高度
// 跟节点的高度就是二叉树的最大深度

func maxDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	res := maxValue(left, right) + 1
	return res
}

func maxValue(i, j int) int {
	if i > j {
		return i
	}
	return j
}
