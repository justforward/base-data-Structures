package main

// 删除倒数第N个节点，就是要指向第N-1个节点
// 快指针需要指到第N+1个节点，才能保证慢指针指到第N-1个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	fast := dummy
	slow := dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
