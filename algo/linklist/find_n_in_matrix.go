package linklist

import "container/heap"

// 查找多个升序矩阵中  第 k 小的元素
// FindKThInMatrix 思路 通过最小堆 合并一条升序链表
func FindKThInMatrix(matrix [][]int, k int) int {
	pq := &MinPriorityQueue{}

	heap.Init(pq)

	for _, ints := range matrix {
		for _, el := range ints {
			heap.Push(pq, el)
		}

	}

	res := -1
	for pq.Len() > 0 && k > 0 {
		res = heap.Pop(pq).(int)
		k--
	}
	return res

}

// 元素存储 []int{数组1元素，数组2元素}
type SumPriorityQueue [][]int

func (pq SumPriorityQueue) Len() int {
	return len(pq)
}

// Less 值小的排在前面
func (pq SumPriorityQueue) Less(i, j int) bool {
	return (pq[i][0] + pq[i][1]) < (pq[j][0] + pq[j][1])
}

// Swap 元素交换
func (pq SumPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *SumPriorityQueue) Push(x any) {
	item := x.([]int)
	*pq = append(*pq, item)
}

func (pq *SumPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

// 查找多个升序矩阵中  第 k 小的元素
// FindKThSumInMatrix 思路 通过最小堆 合并一条升序链表
func FindKThSumInMatrix(u, v []int, k int) [][]int {
	pq := &SumPriorityQueue{}

	heap.Init(pq)

	for _, uElement := range u {
		for _, vElement := range v {
			heap.Push(pq, []int{uElement, vElement})
		}

	}

	res := make([][]int, 0)
	for pq.Len() > 0 && k > 0 {
		element := heap.Pop(pq).([]int)
		res = append(res, element)
		k--
	}
	return res

}

type MinPriorityQueue []int

func (pq MinPriorityQueue) Len() int {
	return len(pq)
}

// Less 值小的排在前面
func (pq MinPriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

// Swap 元素交换
func (pq MinPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinPriorityQueue) Push(x any) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *MinPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	//old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
