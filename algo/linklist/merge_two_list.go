package linklist

func MergeTwoList(l1, l2 *LinkNode) *LinkNode {
	dummy := &LinkNode{Val: -1}

	p := dummy
	p1 := l1
	p2 := l2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}
