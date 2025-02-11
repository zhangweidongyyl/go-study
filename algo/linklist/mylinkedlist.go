package linklist

import "fmt"

// Node  单向链表
type Node struct {
	Val  int
	Next *Node
}
type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewMyLinkedList() *MyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	return &MyLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}
func (mylinkedlist *MyLinkedList) Display() {
	head := mylinkedlist.head
	for p := head.Next; p.Next != nil; p = p.Next {
		fmt.Println("val is ", p.Val)
	}
}

func (list *MyLinkedList) RemoveLast() {
	tmpNode := &Node{}

	head := list.head
	for p := head.Next; p.Next != nil; p = p.Next {
		if p.Next.Next == list.tail {
			tmpNode = p
			return
		}
	}
	tmpNode.Next = list.tail
	list.size--

}

func (list *MyLinkedList) RemoveFirst() {
	head := list.head

	tmpNode := head.Next

	head.Next = tmpNode.Next
	list.size--
}

func (list *MyLinkedList) AddLast(val int) {
	newNode := &Node{Val: val}
	// 只能先找到tail的前一个节点再赋值
	tmpNode := &Node{}
	head := list.head
	for p := head.Next; p.Next != nil; p = p.Next {
		if p.Next == list.tail {
			tmpNode = p
		}
	}
	tmpNode.Next = newNode
	newNode.Next = list.tail

	list.size++
}
func (list *MyLinkedList) AddFirst(val int) {

	newNode := &Node{Val: val}

	tmpNode := list.head.Next

	newNode.Next = tmpNode

	list.head.Next = newNode
	list.size++
}
