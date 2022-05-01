package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root   *Node
	length int
}

func (tree *Tree) insert(newNode Node) *Tree {
	if tree.root == nil {
		tree.root = &newNode
		return tree
	}

	var currentNode *Node
	currentNode = tree.root

	for currentNode != nil {
		if newNode.value < currentNode.value {
			if currentNode.left != nil {
				currentNode = tree.root.left
				continue
			}
			currentNode.left = &newNode
			tree.length++
			return tree
		}

		if newNode.value > currentNode.value {
			if currentNode.right != nil {
				currentNode = tree.root.right
				continue
			}
			currentNode.right = &newNode
			tree.length++
			return tree
		}
	}

	return nil
}

func main() {
	tree := CreateTree()

	tree.insert(Node{value: 4})
	tree.insert(Node{value: 20})
	tree.insert(Node{value: 170})
	tree.insert(Node{value: 15})
	tree.insert(Node{value: 1})
	tree.insert(Node{value: 6})

	fmt.Println(tree.PreOrderTraverse())
}

func CreateTree() *Tree {
	var root Node
	root.value = 9

	tree := new(Tree)
	tree.root = &root
	tree.length = 1

	return tree
}

func (tree *Tree) PreOrderTraverse() (values []int) {
	return tree.root.PreOrderTraverse()
}

func (node *Node) PreOrderTraverse() (values []int) {
	currentNode := node
	values = append(values, currentNode.value)
	if currentNode.left != nil {
		values = append(values, currentNode.left.PreOrderTraverse()...)
	}
	if currentNode.right != nil {
		values = append(values, currentNode.right.PreOrderTraverse()...)
	}
	return
}
