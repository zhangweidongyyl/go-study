package search_ip

import "strings"

type TrieNode struct {
	children map[string]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[string]*TrieNode, 0),
		isEnd:    false,
	}
}

type TrieTree struct {
	root *TrieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{root: NewTrieNode()}
}

func (trieTree *TrieTree) Insert(ip string) {
	ipArr := strings.Split(ip, ".")
	current := trieTree.root
	for _, part := range ipArr {
		if _, exists := current.children[part]; !exists {
			current.children[part] = NewTrieNode()
		}
		current = current.children[part]
	}
	current.isEnd = true
}

func (t *TrieTree) Search(ip string) bool {
	parts := strings.Split(ip, ".")
	current := t.root
	for _, part := range parts {
		if _, exists := current.children[part]; !exists {
			return false
		}
		current = current.children[part]
	}
	return current.isEnd
}
