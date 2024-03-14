package main

/*
链表的归并排序
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwo(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeTwo(list1Node, list2Node *ListNode) *ListNode {
	if list1Node == nil {
		return list2Node
	}
	if list2Node == nil {
		return list1Node
	}
	head := &ListNode{0, nil}
	preNode := head
	list1 := list1Node
	list2 := list2Node
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			preNode.Next = list2
			list2 = list2.Next
		} else {
			preNode.Next = list1
			list1 = list1.Next
		}
		preNode = preNode.Next
	}

	if list1 != nil {
		preNode.Next = list1
	}
	if list2 != nil {
		preNode.Next = list2
	}
	return head.Next
}
