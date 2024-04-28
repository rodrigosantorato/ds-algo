package main

import "fmt"

func main() {
	trie := newNode()

	trie.insert("hello")
	trie.insert("world")
	fmt.Println(trie.search("hello"))   // true
	fmt.Println(trie.search("world"))   // true
	fmt.Println(trie.hasPrefix("hell")) // true
	fmt.Println(trie.hasPrefix("wor"))  // true
	fmt.Println(trie.search("cat"))     // false
	fmt.Println(trie.search("wor"))     // false
	fmt.Println(trie.search(""))        // false
}

type trieNode struct {
	children  map[rune]*trieNode
	endOfWord bool
}

func newNode() *trieNode {
	return &trieNode{
		children: map[rune]*trieNode{},
	}
}

func (n *trieNode) insert(word string) {
	cur := n
	for _, char := range word {
		if cur.children[char] == nil {
			cur.children[char] = newNode()
		}
		cur = cur.children[char]
	}
	cur.endOfWord = true
}

func (n *trieNode) search(word string) bool {
	cur := n
	for _, char := range word {
		if cur.children[char] == nil {
			return false
		}
		cur = cur.children[char]
	}
	return cur.endOfWord
}

func (n *trieNode) hasPrefix(prefix string) bool {
	cur := n
	for _, char := range prefix {
		if cur.children[char] == nil {
			return false
		}
		cur = cur.children[char]
	}
	return true
}
