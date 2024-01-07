package main

// 记住长度即可，等到遍历到最后一个就是
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == length-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
