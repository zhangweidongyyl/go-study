package linklist

// PartitionList 按大小排序，给定相应元素大小，比其小的放前面
func PartitionList(head *LinkNode, val int) *LinkNode {

	dummy1 := &LinkNode{Val: -1}
	dummy2 := &LinkNode{Val: -1}

	p1, p2 := dummy1, dummy2

	p := head

	for p != nil {
		if p.Val < val {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		//p = p.Next
		temp := p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
