package linklist

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func (head *LinkNode) Insert(value int) *LinkNode {
	newHead := &LinkNode{Val: value}
	newHead.Next = head
	return newHead
}
func (head *LinkNode) InsertTail(val int) *LinkNode {
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &LinkNode{Val: val}
	return head
}
func (head *LinkNode) DeleteInIndex(index int) *LinkNode {
	current := head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	//for current.Next != nil {
	//	current = current.Next
	//}
	if current == nil {
		panic("index over length")
	}
	current.Next = current.Next.Next
	return head
}
func (head *LinkNode) InsertInIndex(index, val int) *LinkNode {
	current := head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	//for current.Next != nil {
	//	current = current.Next
	//}
	if current == nil {
		panic("index over length")
	}
	newNode := &LinkNode{Val: val}
	newNode.Next = current.Next
	current.Next = newNode
	return head
}

func CreateLinkNode(arr []int) *LinkNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := &LinkNode{
		Val: arr[0],
	}
	curent := head
	for i := 1; i < len(arr); i++ {
		curent.Next = &LinkNode{Val: arr[i]}
		curent = curent.Next
	}
	return head
}

func PrintLinkNode(node *LinkNode) {
	for p := node; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
