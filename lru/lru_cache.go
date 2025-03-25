package lru

type LRUNode struct {
	key, val   int
	prev, next *LRUNode
}

func NewLRUNode(k, v int) *LRUNode {
	return &LRUNode{
		key:  k,
		val:  v,
		prev: nil,
		next: nil,
	}
}

type DoubleList struct {
	head, tail *LRUNode
	size       int
}

func NewDoubleList() *DoubleList {
	head := &LRUNode{key: 0, val: 0}
	tail := &LRUNode{key: 0, val: 0}
	head.next = tail
	tail.prev = head

	doubleList := &DoubleList{
		head: head,
		tail: tail,
		size: 0,
	}
	return doubleList
}

// AddLast 插入队尾
func (doubleList *DoubleList) AddLast(lruNode *LRUNode) {

	lruNode.prev = doubleList.tail.prev
	lruNode.next = doubleList.tail
	doubleList.tail.prev = lruNode
	doubleList.tail.prev.next = lruNode

	doubleList.size++
}

func (doubleList *DoubleList) Remove(lruNode *LRUNode) {
	lruNode.prev.next = lruNode.next
	lruNode.next.prev = lruNode.prev
	doubleList.size--

}
func (doubleList *DoubleList) RemoveFirst() *LRUNode {
	if doubleList.head.next == doubleList.tail {
		return nil
	}
	firstNode := doubleList.head.next
	doubleList.Remove(firstNode)
	return firstNode
}
func (doubleList *DoubleList) Size() int {
	return doubleList.size
}

type LRUCache struct {
	dataMap map[int]*LRUNode
	cache   DoubleList
	cap     int
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		dataMap: make(map[int]*LRUNode, 0),
		cache:   DoubleList{},
		cap:     cap,
	}
}
func (this *LRUCache) AddRecently(k, v int) {
	lruNode := NewLRUNode(k, v)

	this.cache.AddLast(lruNode)

	this.dataMap[k] = lruNode

}

func (this *LRUCache) makeRecently(k int) {
	lruNode, _ := this.dataMap[k]
	this.cache.Remove(lruNode)
	this.cache.AddLast(lruNode)
}

func (this *LRUCache) deleteKey(k int) {
	lruNode, _ := this.dataMap[k]

	this.cache.Remove(lruNode)

	delete(this.dataMap, k)
}

func (this *LRUCache) removeLastestRecently() {
	deleteNode := this.cache.RemoveFirst()
	deleteKey := deleteNode.key
	this.deleteKey(deleteKey)
}

func (this *LRUCache) Get(k int) int {
	if _, ok := this.dataMap[k]; !ok {
		return -1
	}
	// 将该数据提升为最近使用的
	this.makeRecently(k)
	return this.dataMap[k].val
}
func (this *LRUCache) Put(k, v int) {

	if _, ok := this.dataMap[k]; ok {
		this.deleteKey(k)
		this.AddRecently(k, v)
		return
	}
	if this.cache.Size() == this.cap {
		this.removeLastestRecently()
	}
	this.AddRecently(k, v)

}
