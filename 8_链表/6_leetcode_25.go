package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	root := dummy

	for root != nil {
		begin := pre.Next
		for i := 0; i < k && root != nil; i++ {
			root = root.Next
		}
		if root == nil {
			break
		}
		next := root.Next
		root.Next = nil
		pre.Next = reverseNode(begin)
		begin.Next = next
		pre = begin
		root = begin
	}
	return dummy.Next
}

func reverseNode(head *ListNode) *ListNode {
	pre := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
