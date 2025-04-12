package main

type Node struct {
	children   map[rune]*Node
	isTerminal bool
}
type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{children: map[rune]*Node{}},
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, r := range word {
		childNode, ok := node.children[r]
		if !ok {
			node.children[r] = &Node{children: map[rune]*Node{}}
			childNode = node.children[r]
		}
		node = childNode
	}
	node.isTerminal = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, r := range word {
		childNode, ok := node.children[r]
		if !ok {
			return false
		}
		node = childNode
	}
	return node.isTerminal
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, r := range prefix {
		childNode, ok := node.children[r]
		if !ok {
			return false
		}
		node = childNode
	}
	return true
}
