package myqueue

type MyCycleQueue struct {
	datas *ArrayQueue
	cap   int
}

// ArrayQueue 循环队列
// 注意 记录 头尾节点的位置，以及扩缩容时 头尾节点的变化
type ArrayQueue struct {
	firstIndex int
	lastIndex  int
	size       int
	datas      []int
}

const DefaultArrayQueueCap = 2

func NewDefaultArrayQueue() *ArrayQueue {
	return NewArrayQueue(DefaultArrayQueueCap)
}
func NewArrayQueue(cap int) *ArrayQueue {
	return &ArrayQueue{
		firstIndex: 0,
		lastIndex:  0,
		size:       0,
		datas:      make([]int, cap),
	}
}
func (this *ArrayQueue) Enqueue(data int) {
	if len(this.datas) == this.size {
		this.resize(this.size * 2)
	}
	this.datas[this.lastIndex] = data
	this.lastIndex++
	// 数组满的时候 将lastIndex 置为0 的目的是什么？
	// 这里就是循环 利用空间 最后的 会在此处
	if this.lastIndex == len(this.datas) {
		this.lastIndex = 0
	}
	this.size++

}

// Dequeue 当出队后，前面的内容会空
func (this *ArrayQueue) Dequeue() int {
	oldVal := this.datas[this.firstIndex]
	var zero int
	this.datas[this.firstIndex] = zero
	this.firstIndex++
	if this.firstIndex == len(this.datas) {
		this.firstIndex = 0
	}

	this.size--
	return oldVal
}

// 扩容 ，默认扩大一倍
func (this *ArrayQueue) resize(newCap int) {
	temp := make([]int, newCap)
	// first ----- last
	// --- last    first ---
	for i := 0; i < this.size; i++ {
		temp[i] = this.datas[(this.firstIndex+i)%len(this.datas)]
	}

	this.firstIndex = 0
	this.lastIndex = this.size
	this.datas = temp
}
func (this *ArrayQueue) PeekLast() int {
	if this.lastIndex == 0 {
		return this.datas[len(this.datas)-1]
	}
	return this.datas[this.lastIndex-1]
}
