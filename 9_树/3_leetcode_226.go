package main

// 后序遍历完成，前序遍历也是可以的，
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertTree(root.Left)                         //遍历左节点
	invertTree(root.Right)                        //遍历右节点
	root.Left, root.Right = root.Right, root.Left //交换

	return root

}
