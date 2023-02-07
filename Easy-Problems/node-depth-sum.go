package main

import (
	"fmt"
	"math"
)

type BinaryTreeD struct {
	Value       int
	Left, Right *BinaryTreeD
}
type Level struct {
	Root  *BinaryTreeD
	Depth int
}

// NodeDepths Average case: when the tree is balanced
// O(n) time | O(h) space - where n is the number of nodes in
// the Binary Tree and h is the height of the Binary Tree
func NodeDepths(root *BinaryTreeD) int {
	sumOfDepths := 0
	stack := []Level{{Root: root, Depth: 0}}
	var top Level
	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		node, depth := top.Root, top.Depth
		if node == nil {
			continue
		}
		sumOfDepths += depth
		stack = append(stack, Level{Root: node.Left, Depth: depth + 1})
		stack = append(stack, Level{Root: node.Right, Depth: depth + 1})
	}
	return sumOfDepths
}

func nodeDepthByRecurssion(root *BinaryTreeD) int {
	return nodeDepthsHelper(root, 0)
}

func nodeDepthsHelper(root *BinaryTreeD, depth int) int {
	if root == nil {
		return 0
	}
	return depth + nodeDepthsHelper(root.Left, depth+1) + nodeDepthsHelper(root.Right, depth+1)
}

func sumLastLevel(root *BinaryTreeD, maxLevel int) int {
	if root == nil {
		return 0
	}
	if maxLevel == 1 {
		return root.Value
	}

	return sumLastLevel(root.Left, maxLevel-1) +
		sumLastLevel(root.Right, maxLevel-1)
}

func maxDepth(root *BinaryTreeD) int {
	if root == nil {
		return 0
	}
	return int(1 + math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}

func main() {
	left := BinaryTreeD{5, nil, nil}
	left2 := BinaryTreeD{5, &left, nil}

	right := BinaryTreeD{10, nil, nil}
	right2 := BinaryTreeD{25, nil, &right}

	root := BinaryTreeD{5, &left2, &right2}

	fmt.Println("max depth: ", maxDepth(&root))
	fmt.Println("sum last level: ", sumLastLevel(&root, 3))
	//fmt.Println(NodeDepths(&root))
}
