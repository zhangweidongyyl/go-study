package dfs

import (
	"fmt"
	"math"
)

// 递归遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CreateTree create a tree
//
//	    1
//	2          4
//
// 3     5    6      7
func CreateTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}

	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}

	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	return root
}

var preOrderResult = make([]int, 0)
var middleOrderResult = make([]int, 0)
var postOrderResult = make([]int, 0)

func (this *TreeNode) Traverse() {

	traverse(this)

	fmt.Printf("pre order result is %+v \r\n", preOrderResult)
	fmt.Printf("middle order result is %+v \r\n", middleOrderResult)
	fmt.Printf("post order result is %+v \r\n", postOrderResult)

}
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	//前序
	preOrderResult = append(preOrderResult, root.Val)
	traverse(root.Left)
	// 中序
	middleOrderResult = append(middleOrderResult, root.Val)
	traverse(root.Right)
	// 后序
	postOrderResult = append(postOrderResult, root.Val)
}

// CreateMinDepthTree create a tree
//
//		    1
//		2          4
//
//	            6      7
func CreateMinDepthTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}

	//root.Left.Left = &TreeNode{Val: 3}
	//root.Left.Right = &TreeNode{Val: 5}

	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	return root
}
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDepth := math.MaxInt
	currentDepth := 0
	var traverse1 func(*TreeNode)
	traverse1 = func(root *TreeNode) {
		if root == nil {
			return
		}
		currentDepth++
		if root.Left == nil && root.Right == nil {
			minDepth = min(minDepth, currentDepth)
		}
		//前序
		traverse1(root.Left)
		// 中序
		traverse1(root.Right)
		// 后序
		currentDepth--
	}
	traverse1(root)
	return minDepth
}

func FindAllPath(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	path := make([]int, 0)
	var traverse2 func(*TreeNode)
	traverse2 = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			// important:这里一定是要新构建数组，否则 后续对path的操作会影响加入到res的值
			var newPath []int
			for _, v := range path {
				newPath = append(newPath, v)
			}
			res = append(res, newPath)
		}
		//前序
		traverse2(root.Left)
		// 中序
		traverse2(root.Right)
		// 后序
		path = path[:len(path)-1]
	}
	traverse2(root)
	return res
}
func CountTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftChildCount := CountTree(root.Left)
	rightChildCount := CountTree(root.Right)
	return leftChildCount + rightChildCount + 1
}

func FlipTree(root *TreeNode) {
	if root == nil {
		return
	}
	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp
	FlipTree(root.Left)
	FlipTree(root.Right)
}

type ThreeNodeWithNext struct {
	Val         int
	Left, Right *ThreeNodeWithNext
	Next        *ThreeNodeWithNext
}

func CreateThreeNodeWithNext() *ThreeNodeWithNext {
	root := &ThreeNodeWithNext{Val: 1}
	root.Left = &ThreeNodeWithNext{Val: 2}
	root.Right = &ThreeNodeWithNext{Val: 4}

	root.Left.Left = &ThreeNodeWithNext{Val: 3}
	root.Left.Right = &ThreeNodeWithNext{Val: 5}

	root.Right.Left = &ThreeNodeWithNext{Val: 6}
	root.Right.Right = &ThreeNodeWithNext{Val: 7}

	return root
}
func ConnectThreeNodeWithNext(root *ThreeNodeWithNext) *ThreeNodeWithNext {
	if root == nil {
		return nil
	}
	traverseThreeNodeWithNext(root.Left, root.Right)
	return root
}
func traverseThreeNodeWithNext(left, right *ThreeNodeWithNext) {
	if left == nil || right == nil {
		return
	}
	left.Next = right
	traverseThreeNodeWithNext(left.Left, left.Right)
	traverseThreeNodeWithNext(left.Right, right.Left)
	traverseThreeNodeWithNext(right.Left, right.Right)
}

func BuildMaxBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxVal := 0
	maxValIndex := 0
	for i, num := range nums {
		if maxVal < num {
			maxVal = num
			maxValIndex = i
		}
	}
	root := &TreeNode{
		Val: maxVal,
	}
	root.Left = BuildMaxBinaryTree(nums[:maxValIndex])
	root.Right = BuildMaxBinaryTree(nums[maxValIndex+1:])
	return root
}

// BuildTreeByPreOrder 使用前序遍历和中序遍历 构建二叉树
func BuildTreeByPreOrder(preOrder, middleOrder []int) *TreeNode {
	return build(preOrder, 0, len(preOrder)-1, middleOrder, 0, len(middleOrder)-1)
}
func build(preOrder []int, preStart, preEnd int, middleOrder []int, middleStart, middleEnd int) *TreeNode {

	if preStart > preEnd {
		return nil
	}
	// 前序遍历的 第一个元素 为 root
	rootVal := preOrder[preStart]

	// 找到 根节点 在中序遍历的索引
	rootValIndex := 0
	for i := middleStart; i <= middleEnd; i++ {
		if middleOrder[i] == rootVal {
			rootValIndex = i
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	// 左子树的长度
	leftLength := rootValIndex - middleStart
	root.Left = build(preOrder, preStart+1, preStart+leftLength, middleOrder, middleStart, rootValIndex-1)
	root.Right = build(preOrder, preStart+leftLength+1, preEnd, middleOrder, rootValIndex+1, middleEnd)
	return root
}
