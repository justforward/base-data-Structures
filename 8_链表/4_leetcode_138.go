package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copy特殊的链表，
// 1、针对普通链表直接按照遍历的顺序进行copy
// 2、本题中因为随机指针的存在，当我们拷贝节点时，「当前节点的随机指针指向的节点」可能还没创建，
//
//	利用回溯的思想：让每个节点的copy相互独立，
var cacheNode map[*Node]*Node

func copyRandomList(head *Node) *Node {
	cacheNode = map[*Node]*Node{}
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := cacheNode[node]; ok {
		return n
	}
	newNode := &Node{Val: node.Val}
	cacheNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode

}
