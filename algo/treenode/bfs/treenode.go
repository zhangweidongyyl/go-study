package bfs

import "fmt"

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

// LevelOrderTraverse1 无法知道当前节点所在的层次，无法计算二叉树的深度等，因此该种写法应用场景很少
func LevelOrderTraverse1(root *TreeNode) {
	if root == nil {
		return
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		fmt.Println(current.Val)
		if current.Left != nil {
			q = append(q, current.Left)
		}
		if current.Right != nil {
			q = append(q, current.Right)
		}
	}
}

// LevelOrderTraverse2 层序遍历，并且知晓 当前层
func LevelOrderTraverse2(root *TreeNode) {
	if root == nil {
		return
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	depth := 1

	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			current := q[0]
			q = q[1:]
			fmt.Printf("depte is %d and val is %d \r\n", depth, current.Val)
			if current.Left != nil {
				q = append(q, current.Left)
			}
			if current.Right != nil {
				q = append(q, current.Right)
			}
		}
		depth++

	}
}

type State struct {
	node  *TreeNode
	depth int
}

func LevelOrderTraverse3(root *TreeNode) {
	if root == nil {
		return
	}
	q := make([]State, 0)
	q = append(q, State{
		node:  root,
		depth: 1,
	})
	for len(q) > 0 {
		current := q[0]

		q = q[1:]

		fmt.Printf("current depth is %d and val is %d \r\n", current.depth, current.node.Val)
		if current.node.Left != nil {
			q = append(q, State{node: current.node.Left, depth: current.depth + 1})
		}

		if current.node.Right != nil {
			q = append(q, State{node: current.node.Right, depth: current.depth + 1})
		}
	}
}

// CreateMinDepthTree create a tree
//
//		    1
//		2          4
//
//	         6      7
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

// LevelOrderMinDepth 层序遍历，并且知晓 当前层
func LevelOrderMinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	depth := 1

	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			current := q[0]
			q = q[1:]
			fmt.Printf("depte is %d and val is %d \r\n", depth, current.Val)
			if current.Left == nil && current.Right == nil {
				return depth
			}
			if current.Left != nil {
				q = append(q, current.Left)
			}
			if current.Right != nil {
				q = append(q, current.Right)
			}
		}
		depth++

	}
	return 0
}

type PathState struct {
	node *TreeNode
	path []int
}

func LevelOrderFindAllPath(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	q := make([]PathState, 0)
	q = append(q, PathState{
		node: root,
		path: []int{root.Val},
	})
	for len(q) > 0 {
		current := q[0]

		q = q[1:]
		if current.node.Left == nil && current.node.Right == nil {
			res = append(res, current.path)
		}
		if current.node.Left != nil {
			newPath := append(current.path, current.node.Left.Val)
			q = append(q, PathState{node: current.node.Left, path: newPath})
		}

		if current.node.Right != nil {
			newPath := append(current.path, current.node.Right.Val)
			q = append(q, PathState{node: current.node.Right, path: newPath})
		}
	}
	return res
}
