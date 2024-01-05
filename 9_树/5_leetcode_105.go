package main

// 以 前序遍历为主，按照前序遍历的节点当作，中序遍历的分割线
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	// 从中序中得倒长度，直接从preorder中切割
	mid := len(inorder[:i])
	root.Left = buildTree(preorder[1:mid+1], inorder[:i])
	root.Right = buildTree(preorder[mid+1:], inorder[i+1:])
	return root
}
