package main

import "fmt"

type Node struct {
	Name     string
	Children []*Node
}

//DepthFirstSearch time: O(v + e) | space: O(v) # v:#nodes, e:#edgges in graph
func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}
	return array
}

func main() {
	node1 := Node{"X", nil}
	node2 := Node{"Y", nil}

	array := []*Node{&node1, &node2}

	root := Node{"A", array}

	res := []string{}
	res = root.DepthFirstSearch(res)
	fmt.Println(res)
}
