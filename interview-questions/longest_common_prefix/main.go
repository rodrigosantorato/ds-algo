package main

import "fmt"

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
*/

type node struct {
	children map[rune]*node
}

func newNode() *node {
	return &node{
		children: map[rune]*node{},
	}
}

func (t *node) insert(word string) {
	cur := t
	for _, r := range word {
		if cur.children[r] == nil {
			cur.children[r] = newNode()
		}
		cur = cur.children[r]
	}
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))                         // "fl"
	fmt.Println(longestCommonPrefix([]string{"123cara_flower", "123cara_flow", "123cara_flight"})) // "123cara_fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))                            // ""
}

func longestCommonPrefix(words []string) string {
	res := ""
	trie := &node{children: map[rune]*node{}}
	for _, s := range words {
		trie.insert(s)
	}
	cur := trie
	for i := 0; i < len(words[0]) && cur.children != nil && len(cur.children) == 1; i++ {
		res += string(words[0][i])
		cur = cur.children[rune(words[0][i])]
	}
	return res
}
