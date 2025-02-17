package linklist

import "fmt"

func LongestPalidromeString(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		// 这里取两次是 为了 兼容 str的长度是奇数和偶数
		res1 := getPalidromeString(str, i, i)
		res2 := getPalidromeString(str, i, i+1)
		if len(res) < len(res1) {
			res = res1
		}
		if len(res) < len(res2) {
			res = res2
		}
	}
	return res
}

// 从left往左，right往右 循环 得到一个回文串
func getPalidromeString(str string, left, right int) string {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}
	return str[left+1 : right]
}

func CheckIsPalidromeString(str string) bool {
	if str == "" {
		return false
	}
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func CheckIsPalidromeLink(head *LinkNode) bool {

	left := head
	var right *LinkNode

	var postTraverse func(node *LinkNode)

	var res bool
	postTraverse = func(node *LinkNode) {
		if node == nil {
			return
		}
		postTraverse(node.Next)
		if left.Val != node.Val {
			res = false
		}
		left = left.Next
		fmt.Printf("has enter \r\n")

	}
	right = head
	postTraverse(right)
	return res
}

func Traverse(head *LinkNode) {
	if head == nil {
		return
	}
	fmt.Printf("pre traverse is %d \r\n", head.Val)
	Traverse(head.Next)
	fmt.Println(head.Val)
}

func CheckIsPalidromeLink1(head *LinkNode) bool {

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	left := head
	right := ReverseList(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
