package leetcode

import (
	"container/heap"
)

func getKthMagicNumber(k int) int {
	ugly := -1
	myH := myHeap{}
	heap.Init(&myH)
	heap.Push(&myH, 1)
	for i := 0; i < k; i++ {
		for myH[0] == ugly {
			heap.Pop(&myH)
		}
		ugly = heap.Pop(&myH).(int)
		heap.Push(&myH, 3*ugly)
		heap.Push(&myH, 5*ugly)
		heap.Push(&myH, 7*ugly)
	}
	return ugly
}

// 小根堆
type myHeap []int

func (h myHeap) Len() int {
	return len(h)
}

func (h myHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *myHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
