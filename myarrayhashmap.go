// Copyright 2025 gencozhang. All rights reserved.
// 用数组加强哈希表

package main

import "fmt"

type MyArrayNode struct {
	key string
	val int
}

type MyArrayHashMap struct {
	// m 存储 key 及 key所在的索引位置 key=key在nodes的索引位置
	m map[string]int
	// nodes 存储所有值
	nodes []MyArrayNode
}

func NewMyArrayHashMap() *MyArrayHashMap {
	return &MyArrayHashMap{
		m:     make(map[string]int, 0),
		nodes: make([]MyArrayNode, 0),
	}
}
func (this *MyArrayHashMap) Put(key string, val int) {
	if index, ok := this.m[key]; !ok {
		this.nodes = append(this.nodes, MyArrayNode{
			key: key,
			val: val,
		})
		this.m[key] = len(this.nodes) - 1
		return
	} else {
		this.nodes[index] = MyArrayNode{
			key: key,
			val: val,
		}
	}
}
func (this *MyArrayHashMap) Get(key string) (val int, exist bool) {
	if index, ok := this.m[key]; !ok {
		return 0, false
	} else {
		return this.nodes[index].val, true
	}
}

// Remove 这里就有巧妙的地方，会将要删除的元素挪到最后 nodes就可直接删除最后一个元素
func (this *MyArrayHashMap) Remove(key string) {
	if _, ok := this.m[key]; !ok {
		return
	}
	//要删除元素的index
	index := this.m[key]
	// 要删除的元素
	node := this.nodes[index]

	// 最后一个元素
	lastNode := this.nodes[len(this.nodes)-1]

	this.nodes[index] = lastNode
	this.nodes[len(this.nodes)-1] = node
	this.m[lastNode.key] = index

	this.nodes = this.nodes[:len(this.nodes)-1]
	delete(this.m, key)
}

func main() {
	myArrayHashMap := NewMyArrayHashMap()
	myArrayHashMap.Put("11", 1)
	myArrayHashMap.Put("22", 2)
	myArrayHashMap.Put("33", 3)
	v, exist := myArrayHashMap.Get("22")
	fmt.Printf("v is %+v and exist is %+v \r\n", v, exist)

	myArrayHashMap.Remove("22")

	v1, exist1 := myArrayHashMap.Get("22")
	fmt.Printf("v is %+v and exist is %+v \r\n", v1, exist1)

}
