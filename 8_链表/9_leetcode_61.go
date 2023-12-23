package main

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	length := 0
	root := head
	for root != nil {
		root = root.Next
		length++
	}

}
