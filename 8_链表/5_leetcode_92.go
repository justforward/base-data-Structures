package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// pre 在前一个节点上
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	rightNode := pre
	//
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	leftNode := pre.Next
	curr := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil
	reverse(leftNode)

	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

func reverse(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}
