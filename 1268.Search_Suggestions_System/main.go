package main

import (
	"slices"
)

type Node struct {
	children   map[rune]*Node
	isTerminal bool
	words      []string
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{children: map[rune]*Node{}, words: []string{}},
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, r := range word {
		childNode, ok := node.children[r]
		if !ok {
			node.children[r] = &Node{children: map[rune]*Node{}, words: []string{}}
			childNode = node.children[r]
		}
		node = childNode
		node.words = append(node.words, word)
	}
	node.isTerminal = true
}

func (this *Trie) Search(word string) []string {
	var res []string
	node := this.root
	for _, r := range word {
		childNode, ok := node.children[r]
		if !ok {
			return res
		}
		node = childNode
	}
	if node != nil {
		res = node.words
		slices.Sort(res)
		if len(res) > 3 {
			res = node.words[0:3]
		}

	}
	return res
}

func (this *Trie) InsertMany(words []string) {
	for _, word := range words {
		this.Insert(word)
	}
}

func suggestedProducts(products []string, searchWord string) [][]string {
	var res [][]string
	trie := Constructor()
	trie.InsertMany(products)
	for i := 0; i < len(searchWord); i++ {
		res = append(res, trie.Search(searchWord[0:i+1]))
	}
	return res
}
