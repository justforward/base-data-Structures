package main

// 循环找到链表长度n和长度，计算一下k/n的余数，然后把链表连成环，然后从尾部开始移动n-k，断开环，返回头指针
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	length := 1
	root := head
	for root.Next != nil {
		root = root.Next
		length++
	}
	root.Next = head
	k = k % length
	for i := 0; i < length-k; i++ {
		root = root.Next
	}
	head = root.Next
	root.Next = nil
	return head

}
