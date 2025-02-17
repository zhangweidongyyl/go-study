package linklist

func MoveZero(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}
func DeleteValInPlaceForArray(nums []int, val int) []int {
	if len(nums) == 0 {
		return nums
	}
	fast, slow := 0, 0
	for fast < len(nums) {
		// 如果快指针 不 是要删除的元素  则进行赋值
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return nums[:slow+1]
}

func DeleteRepeatInPlaceForArray(nums []int) []int {

	if len(nums) == 0 {
		return nums
	}
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			// 此时慢指针一定要前移后才重新赋值
			slow++
			// 重新赋值
			nums[slow] = nums[fast]
		}
		fast++
	}
	return nums[:slow+1]

}

func DeleteRepeatInPlaceForLink(head *LinkNode) *LinkNode {
	fast, slow := head, head

	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}
