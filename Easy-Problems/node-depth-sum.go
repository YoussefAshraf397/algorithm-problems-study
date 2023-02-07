package main

import "fmt"

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
	fmt.Println(stack)
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

func NodeDepthByRecurssion(root *BinaryTree) int {
	return nodeDepthsHelper(root, 0)
}

func nodeDepthsHelper(root *BinaryTree, depth int) int {
	if root == nil {
		return 0
	}
	return depth + nodeDepthsHelper(root.Left, depth+1) + nodeDepthsHelper(root.Right, depth+1)
}

func main() {
	left := BinaryTreeD{5, nil, nil}
	left2 := BinaryTreeD{5, &left, nil}

	right := BinaryTreeD{10, nil, nil}
	right2 := BinaryTreeD{10, nil, &right}

	root := BinaryTreeD{5, &left2, &right2}

	fmt.Println(NodeDepths(&root))
}
