package sol

import "container/heap"

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}
func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func isNStraightHand(hand []int, groupSize int) bool {
	nHand := len(hand)
	if groupSize == 1 {
		return true
	}
	if nHand%groupSize != 0 {
		return false
	}
	count := make(map[int]int)
	for _, val := range hand {
		count[val] += 1
	}
	priorityQueue := MinHeap{}
	heap.Init(&priorityQueue)
	for key := range count {
		heap.Push(&priorityQueue, key)
	}
	for priorityQueue.Len() > 0 {
		start := priorityQueue[0]
		end := start + groupSize
		for num := start; num < end; num++ {
			_, ok := count[num]
			if !ok {
				return false
			}
			count[num] -= 1
			if count[num] == 0 {
				heap.Pop(&priorityQueue)
			}
		}
	}
	return true
}
