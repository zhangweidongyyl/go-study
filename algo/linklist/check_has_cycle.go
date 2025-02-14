package linklist

// CheckHasCycle 判断链表中是否成环
// 使用快慢双指针，如果fast slow相遇则有环，否则无环
func CheckHasCycle(head *LinkNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
