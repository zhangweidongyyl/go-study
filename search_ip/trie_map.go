package search_ip

import (
	"fmt"
	"sort"
)

type TrieMapNode struct {
	children map[rune]*TrieMapNode
	val      interface{}
	isEnd    bool
}
type TrieMap struct {
	root *TrieMapNode
}

func NewTrieMapNode() *TrieMapNode {
	return &TrieMapNode{
		children: make(map[rune]*TrieMapNode, 0),
		val:      nil,
		isEnd:    false,
	}
}
func NewTrieMap() *TrieMap {
	return &TrieMap{
		root: NewTrieMapNode(),
	}
}

func (trieMap *TrieMap) Insert(key string, val interface{}) {
	if key == "" {
		panic("insert key cannot nil")
	}
	p := trieMap.root
	for _, c := range key {
		if _, ok := p.children[c]; !ok {
			p.children[c] = NewTrieMapNode()
		}
		p = p.children[c]
	}
	p.val = val
	p.isEnd = true
}
func (trieMap *TrieMap) Search(key string) (interface{}, bool) {
	if key == "" {
		panic("search key cannot nil")
	}
	p := trieMap.root
	for _, c := range key {
		if _, ok := p.children[c]; !ok {
			return nil, false
		}
		p = p.children[c]
	}
	if p.isEnd {
		return p.val, true
	}
	return nil, false
}

func (trieMap *TrieMap) Delete(key string) {
	if key == "" {
		panic("delete key cannot nil")
	}
	p := trieMap.root
	for _, c := range key {
		if _, ok := p.children[c]; !ok {
			fmt.Println("key not exist")
			return
		}
		p = p.children[c]
	}
	p.isEnd = false
	p.val = nil
}

// GetAllWords 按字符顺序排序
func (trieMap *TrieMap) GetAllWords() []string {
	return trieMap.collectWords(trieMap.root, "")
}

// PrefixSearch 前缀查询
func (trieMap *TrieMap) PrefixSearch(prefix string) []string {
	p := trieMap.root
	for _, c := range prefix {
		if _, ok := p.children[c]; !ok {
			return nil
		}
		p = p.children[c]
	}
	return trieMap.collectWords(p, prefix)
}
func (trieMap *TrieMap) collectWords(node *TrieMapNode, prefix string) []string {
	words := make([]string, 0)
	if node.isEnd {
		words = append(words, prefix)
	}
	// 对子节点的字符进行排序
	chars := make([]rune, 0, len(node.children))
	for char := range node.children {
		chars = append(chars, char)
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	// 按字典顺序遍历子节点
	for _, char := range chars {
		child := node.children[char]
		words = append(words, trieMap.collectWords(child, prefix+string(char))...)
	}
	return words
}
