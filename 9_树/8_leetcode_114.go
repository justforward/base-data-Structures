package main

var list []*TreeNode

func flatten(root *TreeNode) {
	list = make([]*TreeNode, 0)
	//前序遍历保持住树的结构
	dfs2(root)
	// 开始将树变成链表形式，left指向nil
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		//
		prev.Left, prev.Right = nil, curr
	}
}

func dfs2(root *TreeNode) {
	if root == nil {
		return
	}
	list = append(list, root)
	dfs2(root.Left)
	dfs2(root.Right)
}
