package linklist

import "fmt"

// GenerateMatrix 按 右 下 左 上 四个点框死矩阵
func GenerateMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	fmt.Printf("matrix is %+v \r\n", matrix)

	// 上边界
	upper_bound := 0
	// 左边界
	left_bound := 0
	//右边界
	right_bound := n - 1
	lower_bound := n - 1
	num := 1
	for num <= n*n {
		if upper_bound <= lower_bound {
			// 在顶部从左向右遍历
			for j := left_bound; j <= right_bound; j++ {
				matrix[upper_bound][j] = num
				num++
			}
			// 上边界下移
			upper_bound++
		}

		if left_bound <= right_bound {
			// 在右侧从上向下遍历
			for i := upper_bound; i <= lower_bound; i++ {
				matrix[i][right_bound] = num
				num++
			}
			// 右边界左移
			right_bound--
		}

		if upper_bound <= lower_bound {
			// 在底部从右向左遍历
			for j := right_bound; j >= left_bound; j-- {
				matrix[lower_bound][j] = num
				num++
			}
			// 下边界上移
			lower_bound--
		}

		if left_bound <= right_bound {
			// 在左侧从下向上遍历
			for i := lower_bound; i >= upper_bound; i-- {
				matrix[i][left_bound] = num
				num++
			}
			// 左边界右移
			left_bound++
		}
	}
	return matrix
}

// RotateMatrix 二维数组的翻转，矩阵的翻转
func RotateMatrix(matrix [][]int) [][]int {
	// 先进行 90度翻转
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 后进行一维数组翻转
	for i := 0; i < len(matrix); i++ {
		matrix[i] = reverse(matrix[i])
	}
	return matrix
}
func reverse(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		t := arr[left]
		arr[left] = arr[right]
		arr[right] = t
		left++
		right--
	}
	return arr
}
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
