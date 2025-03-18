package lru

type LRUNode struct {
}

type LRUCache struct {
	dataMap map[int]*LRUNode
}
