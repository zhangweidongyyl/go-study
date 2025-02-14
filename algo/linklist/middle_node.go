package linklist

// MiddleNode 使用快慢双指针 找到链表的重点 fast每次走两步 slow 每次走一步，当fast走到末尾时 end刚好走到中带你
func MiddleNode(head *LinkNode) *LinkNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
