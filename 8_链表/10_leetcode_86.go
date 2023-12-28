package main

func partition(head *ListNode, x int) *ListNode {
	root := head
	dummy1 := &ListNode{Val: -1}
	head1 := dummy1
	dummy2 := &ListNode{
		Val: -2,
	}
	head2 := dummy2

	for root != nil {
		if root.Val < x {
			head1.Next = root
			head1 = head1.Next
		} else {
			head2.Next = root
			head2 = head2.Next
		}
		root = root.Next
	}

	head2.Next = nil // 遍历结束后，我们将 large next指针置空，这是因为当前节点复用的是原链表的节点，而其 next 指针可能指向一个小于 xxx 的节点，我们需要切断这个引用。
	head1.Next = head2.Next
	return dummy1.Next
}
