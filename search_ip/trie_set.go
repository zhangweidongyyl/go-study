package search_ip

import "fmt"

type TrieSetNode struct {
	children map[rune]*TrieSetNode
	isEnd    bool
}
type TrieSet struct {
	root *TrieSetNode
}

func NewTrieSetNode() *TrieSetNode {
	return &TrieSetNode{
		children: make(map[rune]*TrieSetNode, 0),
		isEnd:    false,
	}
}
func NewTrieSet() *TrieSet {
	return &TrieSet{
		root: NewTrieSetNode(),
	}
}

func (trieSet *TrieSet) Insert(key string) {
	if key == "" {
		panic("insert key cannot nil")
	}
	p := trieSet.root
	for _, c := range key {
		if _, ok := p.children[c]; !ok {
			p.children[c] = NewTrieSetNode()
		}
		p = p.children[c]
	}
	p.isEnd = true
}
func (trieSet *TrieSet) Contains(key string) bool {
	if key == "" {
		panic("contains key cannot nil")
	}
	p := trieSet.root
	for _, c := range key {
		if _, ok := p.children[c]; !ok {
			return false
		}
		p = p.children[c]
	}
	return p.isEnd
}

func (trieSet *TrieSet) Delete(key string) {
	if key == "" {
		panic("delete key cannot nil")
	}
	p := trieSet.root
	for _, c := range key {
		if _, ok := p.children[c]; !ok {
			fmt.Println("key not exist")
			return
		}
		p = p.children[c]
	}
	p.isEnd = false
}
