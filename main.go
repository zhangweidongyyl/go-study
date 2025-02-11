package main

import (
	"fmt"
	"study.com/study/algo/treenode/bfs"
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

	//root := dfs.CreateMinDepthTree()
	//res := dfs.FindAllPath(root)
	//for index, path := range res {
	//	fmt.Printf("path index is %d and path is %+v \r\n", index, path)
	//}

	root := bfs.CreateMinDepthTree()
	res := bfs.LevelOrderFindAllPath(root)
	for index, path := range res {
		fmt.Printf("path index is %d and path is %+v \r\n", index, path)
	}
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
