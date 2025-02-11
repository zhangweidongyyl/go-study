// Copyright 2025 gencozhang. All rights reserved.
// 用双链表加强哈希表

package main

import "fmt"

type Node struct {
	key  string
	val  int
	prev *Node
	next *Node
}

type MyLinkedHashMap struct {
	head *Node
	tail *Node
	m    map[string]*Node
}

func New() *MyLinkedHashMap {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &MyLinkedHashMap{
		head: head,
		tail: tail,
		m:    make(map[string]*Node, 0),
	}
}

func (this *MyLinkedHashMap) Put(key string, val int) {
	if _, ok := this.m[key]; !ok {

		this.addLastNode(&Node{
			key:  key,
			val:  val,
			prev: nil,
			next: nil,
		})
		return
	}
	this.m[key] = &Node{
		key:  key,
		val:  val,
		prev: nil,
		next: nil,
	}
}
func (this *MyLinkedHashMap) Get(key string) (val int, exist bool) {

	if node, ok := this.m[key]; ok {
		return node.val, ok
	}
	return 0, false
}
func (this *MyLinkedHashMap) Keys() []string {
	keys := make([]string, 0)
	for p := this.head.next; p != this.tail; p = p.next {
		keys = append(keys, p.key)
	}
	return keys
}

// addLastNode插入到最后
func (this *MyLinkedHashMap) addLastNode(addNode *Node) {
	temp := this.tail.prev

	// 双向连接 addNode和tail
	addNode.next = this.tail
	this.tail.prev = addNode

	// 双向连接 temp 和addNode
	addNode.prev = temp
	temp.next = addNode

	// 新增加的node添加到map中
	this.m[addNode.key] = addNode

}

func main11() {
	linkedMap := New()

	linkedMap.Put("aa", 1)
	linkedMap.Put("ddc", 2)
	v, ok := linkedMap.Get("aa")
	fmt.Printf("%+v is exist %+v \r\n", v, ok)
	fmt.Printf("%+v \r\n", linkedMap.Keys())
}
