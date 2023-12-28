package main

func isSymmetric(root *TreeNode) bool {
	return is(root, root)
}

func is(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val {
		return is(root1.Left, root2.Right) && is(root1.Right, root2.Left)
	} else {
		return false
	}
}
