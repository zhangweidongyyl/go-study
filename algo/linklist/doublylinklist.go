package linklist

import "fmt"

type DoublyLinkNode struct {
	Val        int
	Prev, Next *DoublyLinkNode
}

func NewDoublyLinkNode(val int) *DoublyLinkNode {
	return &DoublyLinkNode{
		Val: val,
	}
}

func CreateDoublyLinkNode(arr []int) *DoublyLinkNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := &DoublyLinkNode{
		Val: arr[0],
	}
	current := head
	for i := 1; i < len(arr); i++ {
		newNode := &DoublyLinkNode{Val: arr[i]}
		current.Next = newNode
		newNode.Prev = current
		current = current.Next
	}
	return head
}

func PrintDoublyLinkNode(node *DoublyLinkNode) {
	fmt.Println("print from head")
	tail := &DoublyLinkNode{}
	for p := node; p != nil; p = p.Next {
		fmt.Println(p.Val)
		tail = p
	}
	fmt.Println("tail val:", tail.Val)
	dummyHead := &DoublyLinkNode{}
	for p := tail; p != nil; p = p.Prev {
		fmt.Println(p.Val)
		dummyHead = p
	}
	fmt.Println("head val:", dummyHead.Val)
}
func InsertInHead(node *DoublyLinkNode, val int) *DoublyLinkNode {
	newHead := NewDoublyLinkNode(val)
	newHead.Next = node
	node.Prev = newHead
	return newHead
}

func InsertInTail(node *DoublyLinkNode, val int) (*DoublyLinkNode, *DoublyLinkNode) {
	tail := node
	for tail.Next != nil {
		tail = tail.Next
	}
	newTail := NewDoublyLinkNode(val)
	tail.Next = newTail
	newTail.Prev = tail

	return node, newTail
}
