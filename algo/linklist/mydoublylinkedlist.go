package linklist

import "fmt"

type DoublyNode struct {
	Val  int
	Prev *DoublyNode
	Next *DoublyNode
}
type MyDoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
	size int
}

func NewDoublyLinkedList() *MyDoublyLinkedList {
	head := &DoublyNode{}
	tail := &DoublyNode{}

	head.Next = tail
	tail.Prev = head
	return &MyDoublyLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (list *MyDoublyLinkedList) AddFirst(val int) {
	newNode := &DoublyNode{
		Val:  val,
		Prev: nil,
		Next: nil,
	}
	tmpNode := list.head.Next

	tmpNode.Prev = newNode
	newNode.Next = tmpNode

	list.head.Next = newNode

	list.size++

}

func (list *MyDoublyLinkedList) Display() {
	for p := list.head; p.Next != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
