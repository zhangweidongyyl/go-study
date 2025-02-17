package linklist

func ReverseList1(head *LinkNode) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// ReverseBetween 先找到第m-1 个节点  然后
func ReverseBetween(head *LinkNode, m, n int) *LinkNode {

	if head == nil || head.Next == nil {
		return head
	}
	p := head
	for i := 1; i < m-1; i++ {
		p = p.Next
	}
	p.Next = ReverseN(p.Next, n-(m-1))
	return head
}

func ReverseKGroup(head *LinkNode, k int) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 先找到前k个元素的 列表
	a, b := head, head

	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := ReverseN(a, k)
	a.Next = ReverseKGroup(b, k)
	return newHead
}

func ReverseN(head *LinkNode, n int) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *LinkNode
	pre, current, next := nil, head, head.Next
	for n > 0 {
		current.Next = pre
		pre = current
		current = next
		if next != nil {
			next = next.Next
		}

		n--
	}
	head.Next = current
	return pre
}

// ReverseList 链表的 反转
func ReverseList(head *LinkNode) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *LinkNode
	pre, current, next := nil, head, head.Next

	for current != nil {
		current.Next = pre
		// 全部往后移动
		pre = current
		current = next
		if next != nil {
			next = next.Next
		}
	}
	return pre
}
