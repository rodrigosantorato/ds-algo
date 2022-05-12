package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root   *Node
	length int
}

func main() {
	tree := CreateTree(9)

	tree.insert(Node{value: 4})
	tree.insert(Node{value: 20})
	tree.insert(Node{value: 170})
	tree.insert(Node{value: 15})
	tree.insert(Node{value: 1})
	tree.insert(Node{value: 6})
	//fmt.Println(tree.lookup(6))
	//fmt.Println(tree.lookup(170))
	//fmt.Println(tree.lookup(1))
	//fmt.Println(tree.lookup(4))
	//fmt.Println(tree.lookup(190))
	//fmt.Println(tree.breadthFirstSearch())
	fmt.Println(tree.DFSInOrder(tree.root, &[]int{}))

	//fmt.Println(tree.PreOrderTraverse())

	//jsonTree, _ := json.Marshal(tree)
	//fmt.Printf("JSON Print - \n%s\n", string(jsonTree))
	//
	//yamlTree, _ := yaml.Marshal(tree)
	//fmt.Printf("JSON Print - \n%s\n", string(yamlTree))

	//fmt.Println(render.Render(tree))

	//fmt.Printf("%+v\n", tree)
}

func CreateTree(rootValue int) *Tree {
	var root Node
	root.value = rootValue

	tree := new(Tree)
	tree.root = &root
	tree.length = 1

	return tree
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

		if currentNode.right != nil {
			currentNode = tree.root.right
			continue
		}
		currentNode.right = &newNode
		tree.length++
		return tree
	}

	return nil
}

func (tree *Tree) lookup(value int) *Node {
	currentNode := tree.root

	for currentNode != nil {
		if value == currentNode.value {
			return currentNode
		}

		if value < currentNode.value {
			currentNode = currentNode.left
			continue
		}
		currentNode = currentNode.right
	}

	return nil
}

func (tree *Tree) breadthFirstSearch() []int {
	currentNode := tree.root
	var arr []int
	queue := []*Node{currentNode}

	for len(queue) > 0 {
		currentNode, queue = queue[0], queue[1:]
		arr = append(arr, currentNode.value)
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}
	return arr
}

func (tree *Tree) DFSInOrder(node *Node, arr *[]int) []int {
	if node.left != nil {
		tree.DFSInOrder(node.left, arr)
	}
	*arr = append(*arr, node.value)
	if node.right != nil {
		tree.DFSInOrder(node.right, arr)
	}
	return *arr
}

//func (tree *Tree) PreOrderTraverse() (values []int) {
//	return tree.root.PreOrderTraverse()
//}
//
//func (node *Node) PreOrderTraverse() (values []int) {
//	currentNode := node
//	values = append(values, currentNode.value)
//	if currentNode.left != nil {
//		values = append(values, currentNode.left.PreOrderTraverse()...)
//	}
//	if currentNode.right != nil {
//		values = append(values, currentNode.right.PreOrderTraverse()...)
//	}
//	return
//}
