package main

// 所有数值都是唯一的
// p q为 不同节点且均在给定的二叉树中
//

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	// 这个是要把公共祖先节点返回到上层去，那一层不为空说明，需要返回哪一层
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
