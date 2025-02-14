package linklist

func FindFromEnd(head *LinkNode, k int) *LinkNode {
	p1 := head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}

	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

// RemoveFromEnd 删除倒数第n个节点
func RemoveFromEnd(head *LinkNode, n int) *LinkNode {
	dummy := LinkNode{Val: -1}
	dummy.Next = head

	preDelNode := FindFromEnd(head, n+1)
	preDelNode.Next = preDelNode.Next.Next

	return dummy.Next
}
