package main

// 中序和后序遍历
// 后序遍历的数组最后一个元素代表的就是根节点
// 利用已知的跟节点信息在中序遍历的数组中找到根节点的下标
// 然后根据其将中序遍历的数组分成左右两部分，左边部分即左子树，右边部分为右子树，
func buildTree2(inorder []int, postorder []int) *TreeNode {
	idxMap := map[int]int{}
	for i, v := range inorder {
		idxMap[i] = v
	}

	return &TreeNode{}
}
