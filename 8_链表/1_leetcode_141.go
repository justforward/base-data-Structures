package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	mapTmp := make(map[*ListNode]struct{})
	for head.Next != nil {
		if _, exist := mapTmp[head]; exist {
			return true
		} else {
			mapTmp[head] = struct{}{}
		}
		head = head.Next
	}
	return false
}
