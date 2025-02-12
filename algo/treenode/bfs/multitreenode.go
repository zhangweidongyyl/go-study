package bfs

import "fmt"

type Node struct {
	val      int
	children []*Node
}

// CreateNTree create multi tree
//
//	    1
//	2   3   4
//
// 5  6  7
func CreateNTree() *Node {
	root := &Node{val: 1}
	child11 := &Node{val: 5}
	child12 := &Node{val: 6}
	child13 := &Node{val: 7}

	child1 := &Node{val: 2}
	child1.children = make([]*Node, 0)
	child1.children = append(child1.children, child11, child12, child13)

	child2 := &Node{val: 3}
	child3 := &Node{val: 4}
	root.children = make([]*Node, 0)
	root.children = append(root.children, child1, child2, child3)
	return root
}

func NtreeLevelOrderTraverse1(root *Node) {
	if root == nil {
		return
	}
	q := make([]*Node, 0)
	q = append(q, root)
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		fmt.Printf("ntree level order traverse1 :%d \r\n", current.val)
		for _, child := range current.children {
			q = append(q, child)
		}
	}
}

func NtreeLevelOrderTraverse2(root *Node) {
	if root == nil {
		return
	}
	q := make([]*Node, 0)
	q = append(q, root)
	depth := 1
	for len(q) > 0 {

		sz := len(q)
		for i := 0; i < sz; i++ {
			current := q[0]
			q = q[1:]
			fmt.Printf("depth is %d ntree level order traverse2 :%d \r\n", depth, current.val)
			for _, child := range current.children {
				q = append(q, child)
			}
		}

		depth++
	}
}

type NtreeState struct {
	node  *Node
	depth int
}

func NtreeLevelOrderTraverse3(root *Node) {
	if root == nil {
		return
	}
	q := make([]NtreeState, 0)
	q = append(q, NtreeState{
		node:  root,
		depth: 1,
	})
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		fmt.Printf("depth is %d ntree level order traverse3 :%d \r\n", current.depth, current.node.val)
		for _, child := range current.node.children {
			q = append(q, NtreeState{
				node:  child,
				depth: current.depth + 1,
			})
		}
	}
}
