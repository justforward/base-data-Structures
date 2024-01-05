package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	tmp := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		if tmp == 0 {
			return true
		}
	}
	return hasPathSum(root.Right, tmp) || hasPathSum(root.Left, tmp)
}
