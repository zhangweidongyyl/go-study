package myqueue

import "container/list"

type RecentNumber struct {
	queue []int
}

func NewRecentNumber() *RecentNumber {
	return &RecentNumber{queue: make([]int, 0)}
}
func (this *RecentNumber) Ping(t int) int {
	this.queue = append(this.queue, t)
	for this.queue[0] < t-3000 {
		// t 是递增的，所以可以从队头删除 3000 毫秒之前的请求
		this.queue = this.queue[1:]
	}
	return len(this.queue)

}

type MyQueue1 struct {
	list *list.List
}

func NewMyQueue1() MyQueue1 {
	return MyQueue1{
		list: list.New(),
	}
}
func (this *MyQueue1) Push(element int) {
	this.list.PushFront(element)
}
func (this *MyQueue1) Peek() int {
	return this.list.Front().Value.(int)
}

type MyQueue struct {
	tail, head []int
}

func (this *MyQueue) Push(element int) {
	if len(this.tail) == 0 {
		this.tail = make([]int, 0)
	}
	this.tail = append(this.tail, element)
}

// Peek 从对头读元素
func (this *MyQueue) Peek() int {
	if len(this.head) == 0 {
		this.head = make([]int, 0)
	}
	// push的时候会将元素 全部加入到 head 切片里
	for len(this.tail) > 0 {
		this.head = append(this.head, this.tail[len(this.tail)-1])
		this.tail = this.tail[:len(this.tail)-1]
	}
	return this.head[len(this.head)-1]
}
func (this *MyQueue) Pop() int {
	// 有元素 保证会挪到this.head 数组里
	this.Peek()
	res := this.head[len(this.head)-1]
	this.head = this.head[:len(this.head)-1]
	return res
}
func (this *MyQueue) Empty() bool {
	return len(this.head) == 0 && len(this.tail) == 0
}
