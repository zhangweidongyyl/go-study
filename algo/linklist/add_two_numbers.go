package linklist

// AddTwoNumbers  逆序两个队列
func AddTwoNumbers(l1, l2 *LinkNode) *LinkNode {
	dummy := &LinkNode{Val: -1}
	p := dummy

	p1, p2 := l1, l2

	// 每次加的进位
	carry := 0
	for p1 != nil || p2 != nil || carry > 0 {
		val := carry
		if p1 != nil {
			val += p1.Val
			p1 = p1.Next
		}

		if p2 != nil {
			val += p2.Val
			p2 = p2.Next
		}
		// val 加完后 可能超过10 此时需要进行处理 是否往下一步的进位
		carry = val / 10

		val = val % 10

		p.Next = &LinkNode{
			Val: val,
		}
		p = p.Next
	}

	return dummy.Next
}

func AddTwoNumbers2(l1, l2 *LinkNode) *LinkNode {
	p1, p2 := l1, l2
	stk1 := make([]int, 0)
	for p1 != nil {
		stk1 = append(stk1, p1.Val)
		p1 = p1.Next
	}
	stk2 := make([]int, 0)
	for p2 != nil {
		stk2 = append(stk2, p2.Val)
		p2 = p2.Next
	}

	carry := 0
	dummy := &LinkNode{Val: -1}
	p := dummy

	for len(stk1) > 0 || len(stk2) > 0 || carry > 0 {
		// 加上上次的进位
		val := carry
		if len(stk1) > 0 {
			val += stk1[len(stk1)-1]
			stk1 = stk1[:len(stk1)-1]
		}

		if len(stk2) > 0 {
			val += stk2[len(stk2)-1]
			stk2 = stk2[:len(stk2)-1]
		}
		// 处理下一位的进位
		carry = val / 10

		val = val % 10

		p.Next = &LinkNode{Val: val}
		p = p.Next
	}
	return dummy.Next
}
