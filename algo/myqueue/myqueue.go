package myqueue

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
	for len(this.tail) > 0 {
		this.head = append(this.head, this.tail[len(this.tail)-1])
		this.tail = this.tail[:len(this.tail)-1]
	}
	return this.head[len(this.head)-1]
}
