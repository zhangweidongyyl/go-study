package linklist

import (
	"container/heap"
	"fmt"
)

type PriorityQueue []*LinkNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less 值小的排在前面
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

// Swap 元素交换
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*LinkNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func MergeKList(list []*LinkNode) *LinkNode {

	pq := &PriorityQueue{}
	heap.Init(pq)
	// push in head
	for _, node := range list {
		if node != nil {
			heap.Push(pq, node)
		}
	}

	resultDummy := &LinkNode{Val: -1}
	p := resultDummy
	for pq.Len() > 0 {
		// 此时取到的node肯定是所有给到的 node里最小的
		node := heap.Pop(pq).(*LinkNode)
		fmt.Printf("node is %+v \r\n", node)
		p.Next = node
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
		p = p.Next
	}
	return resultDummy.Next
}
