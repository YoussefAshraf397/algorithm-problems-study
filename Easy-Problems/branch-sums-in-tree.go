package main

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// BranchSums time: o(n) | space: o(n) where n is number of nodes in BT
func BranchSums(root *BinaryTree) []int {
	sums := []int{}
	calculateBranchSums(root, 0, &sums)
	return sums
}

func calculateBranchSums(node *BinaryTree, runningSum int, sums *[]int) {
	if node == nil {
		return
	}
	runningSum += node.Value
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, runningSum)
		return
	}
	calculateBranchSums(node.Left, runningSum, sums)
	calculateBranchSums(node.Right, runningSum, sums)
}

func main() {
	left := BinaryTree{5, nil, nil}
	right := BinaryTree{15, nil, nil}
	x := BinaryTree{20, &left, &right}
	fmt.Println(BranchSums(&x))
}
