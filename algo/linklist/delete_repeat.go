package linklist

// DeleteRepeat 删除升序 队列里重复的元素，返回非重复的元素
func DeleteRepeat(head *LinkNode) *LinkNode {

	p := head

	dummyRepeat := &LinkNode{Val: -1}
	dummyUniq := &LinkNode{Val: -1}

	pRepeat := dummyRepeat
	pUniq := dummyUniq

	for p.Next != nil {
		// 把重复的 放到pRepeat后 非重复 的放到pUniq后
		// 连着两个重复，或和pRepeat的重复
		if p.Val == p.Next.Val || p.Val == pRepeat.Val {
			pRepeat.Next = p
			pRepeat = pRepeat.Next
		} else {
			pUniq.Next = p
			pUniq = pUniq.Next

		}

		p = p.Next
	}

	return dummyUniq.Next
}

func DeleteRepeat2(head *LinkNode) *LinkNode {

	dummyUniq := &LinkNode{Val: -1}

	pUniq := dummyUniq
	p := head

	for p != nil {

		if p.Next != nil && (p.Val == p.Next.Val) {

			for p.Next != nil && (p.Val == p.Next.Val) {
				p = p.Next
			}

			p = p.Next
			if p == nil {
				pUniq.Next = nil
			}
		} else {
			pUniq.Next = p
			p = p.Next
			pUniq = pUniq.Next
		}
	}
	return dummyUniq.Next
}
