package main

import (
	"fmt"
	"study.com/study/algo/linklist"
	"sync"
)

var wg sync.WaitGroup

func goroutine1(i int) {
	fmt.Printf("hello goroutine!:%d \n", i)
	wg.Done()
}

type UserInfoStruct struct {
	uid  uint32
	name string
}

func (userinfo UserInfoStruct) getusername() (string, error) {
	if userinfo.uid >= 100 {
		return "bigusername", nil
	} else {
		return "smallusername", nil
	}
}
func (userinfo UserInfoStruct) setusername() (bool, error) {
	if userinfo.uid >= 100 {
		return true, nil
	} else {
		return false, nil
	}
}
func main() {
	// linklist相关
	//1、merge two link list
	//
	//l1 := linklist.CreateLinkNode([]int{1, 3, 5})
	//l2 := linklist.CreateLinkNode([]int{2, 6, 9, 14, 10})
	//l3 := linklist.MergeTwoList(l1, l2)
	////linklist.PrintLinkNode(l3)
	//
	//// 2 、 partition link list
	//l4 := linklist.PartitionList(l3, 5)
	//linklist.PrintLinkNode(l4)

	//l1 := linklist.CreateLinkNode([]int{1, 3, 5})
	//l2 := linklist.CreateLinkNode([]int{2, 6, 9, 14, 8})
	//
	//res := linklist.MiddleNode(l2)
	//linklist.PrintLinkNode(res)
	//
	//l2Repeat := linklist.CreateLinkNode([]int{2, 2, 3, 6, 6, 8})
	//
	//resUniq := linklist.DeleteRepeat2(l2Repeat)
	//linklist.PrintLinkNode(resUniq)
	//res := linklist.RemoveFromEnd(l2, 2)
	//linklist.PrintLinkNode(res)

	//l1 := linklist.CreateLinkNode([]int{1, 3, 5})
	//l2 := linklist.CreateLinkNode([]int{2, 6, 9, 14})
	//list := make([]*linklist.LinkNode, 0)
	//list = append(list, l1, l2)
	//merge k list 使用优先级队列 合并k个升序队列
	//res := linklist.MergeKList(list)
	////
	//linklist.PrintLinkNode(res)
	//
	////fmt.Printf("ugly number %d is %d \r\n", 14, linklist.UglyNumberOfN(14))
	//fmt.Printf("res is %+v \r\n", linklist.IsUglyNumber(6))

	//matrix := make([][]int, 0)
	//e1 := []int{1, 2, 3, 4}
	//e2 := []int{4, 5, 6, 7}
	//matrix = append(matrix, e1, e2)
	//k := 7
	//
	//res := linklist.FindKThInMatrix(matrix, k)
	//fmt.Printf("the %d min number in matrix is %d", k, res)

	//e1 := []int{1, 2, 7}
	//e2 := []int{2, 4, 6}
	//k := 3
	//
	//res := linklist.FindKThSumInMatrix(e1, e2, k)
	//fmt.Printf("the %d min number in matrix is %d", k, res)

	e1 := []int{1, 2, 7}
	l1 := linklist.CreateLinkNode(e1)
	e2 := []int{2, 4, 6}
	l2 := linklist.CreateLinkNode(e2)
	res := linklist.AddTwoNumbers2(l1, l2)
	linklist.PrintLinkNode(res)

	//treenode.PrintPq()
	//ntreeroot := dfs.CreateNTree()
	//dfs.TraverseNTree(ntreeroot)

	//ntreeroot := bfs.CreateNTree()
	//bfs.NtreeLevelOrderTraverse3(ntreeroot)

	//root := dfs.CreateMinDepthTree()
	//res := dfs.FindAllPath(root)
	//for index, path := range res {
	//	fmt.Printf("path index is %d and path is %+v \r\n", index, path)
	//}

	//root := bfs.CreateMinDepthTree()
	//res := bfs.LevelOrderFindAllPath(root)
	//for index, path := range res {
	//	fmt.Printf("path index is %d and path is %+v \r\n", index, path)
	//}
	//minDepth := dfs.MinDepth(root)
	//fmt.Println(minDepth)
	//root := dfs.CreateTree()
	//root.Traverse()
	//root := bfs.CreateTree()
	//bfs.LevelOrderTraverse1(root)
	//bfs.LevelOrderTraverse2(root)
	//bfs.LevelOrderTraverse3(root)

	//root := bfs.CreateMinDepthTree()
	//fmt.Println(bfs.LevelOrderMinDepth(root))

}
