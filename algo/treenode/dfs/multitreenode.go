package dfs

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

var ntreePreOrderResult []int
var ntreePostOrderResult []int

func TraverseNTree(root *Node) {
	traverseNTree(root)
	fmt.Printf("ntree pre result is %+v \r\n", ntreePreOrderResult)
	fmt.Printf("ntree post result is %+v \r\n", ntreePostOrderResult)
}
func traverseNTree(root *Node) {
	if root == nil {
		return
	}
	// 在这里前序
	ntreePreOrderResult = append(ntreePreOrderResult, root.val)
	if len(root.children) > 0 {
		for _, child := range root.children {
			traverseNTree(child)
		}

	}
	// 在这里写后序
	ntreePostOrderResult = append(ntreePostOrderResult, root.val)
}
