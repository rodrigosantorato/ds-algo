package main

import (
	"fmt"
	"github.com/gdexlab/go-render/render"
)

func main() {
	myGraph := NewGraph()

	myGraph.addVertex(0)
	myGraph.addVertex(1)
	myGraph.addVertex(2)
	myGraph.addVertex(3)
	myGraph.addVertex(4)
	myGraph.addVertex(5)
	myGraph.addVertex(6)
	myGraph.addEdge(3, 1)
	myGraph.addEdge(3, 4)
	myGraph.addEdge(4, 2)
	myGraph.addEdge(4, 5)
	myGraph.addEdge(1, 2)
	myGraph.addEdge(0, 1)
	myGraph.addEdge(2, 0)
	myGraph.addEdge(6, 5)

	fmt.Println(render.Render(myGraph))
}

type Graph struct {
	numberOfNodes int
	adjacentList  map[int][]int
}

func NewGraph() *Graph {
	newGraph := new(Graph)
	newGraph.numberOfNodes = 0
	newGraph.adjacentList = make(map[int][]int)

	return newGraph
}

func (graph *Graph) addVertex(vertex int) {
	graph.adjacentList[vertex] = []int{}
	graph.numberOfNodes++
}

func (graph *Graph) addEdge(node1 int, node2 int) {
	graph.adjacentList[node1] = append(graph.adjacentList[node1], node2)
	graph.adjacentList[node2] = append(graph.adjacentList[node2], node1)
}
