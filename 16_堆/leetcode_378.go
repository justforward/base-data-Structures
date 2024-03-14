package main

import "container/heap"

/*
有序数组中的，第k小的元素
从每一行的第一个开始放入到堆中，一直放入k个元素。
*/
func kthSmallest(matrix [][]int, k int) int {
	h := &IHeap{}
	for i := 0; i < len(matrix); i++ {
		heap.Push(h, [3]int{matrix[i][0], i, 0})
	}

	for i := 0; i < k-1; i++ {
		now := heap.Pop(h).([3]int)
		if now[2] != len(matrix)-1 {
			nowXIndex := now[1]
			nowYIndex := now[2]
			heap.Push(h, [3]int{matrix[nowXIndex][nowYIndex+1], nowXIndex, nowYIndex + 1})
		}
	}
	return heap.Pop(h).([3]int)[0]
}

type IHeap [][3]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
