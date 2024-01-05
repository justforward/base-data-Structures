package main

// 使用二叉树节点记录公式，如果是满二叉树 节点个数为2^高度
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ldepth := 0
	cur := root.Left
	for cur != nil {
		ldepth++
		cur = cur.Left
	}
	rdepth := 0
	cur = root.Right
	for cur != nil {
		rdepth++
		cur = cur.Right
	}

	if ldepth == rdepth {
		return 2<<ldepth - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
