package main

import (
	"fmt"
	"math"
)

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// FindClosestValue time: o(log(n)) | space: o(log(n))
// FindClosestValue WORST time: o(n) | space: o(n)
func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosestValue(target, math.MaxInt32)
}

func (tree *BST) findClosestValue(target, closest int) int {
	if absdiff(target, closest) > absdiff(target, tree.Value) {
		closest = tree.Value
	}
	if target < tree.Value && tree.Left != nil {
		return tree.Left.findClosestValue(target, closest)
	} else if target > tree.Value && tree.Right != nil {
		return tree.Right.findClosestValue(target, closest)
	}
	return closest
}

// FindClosestValueOptimize time: o(log(n)) | space: o(1)
// FindClosestValueOptimize WORST time: o(n) | space: o(1)
func (tree *BST) FindClosestValueOptimize(target int) int {
	return tree.findClosestValueOptimize(target, math.MaxInt32)
}

func (tree *BST) findClosestValueOptimize(target, closest int) int {
	currentNode := tree
	for currentNode != nil {
		if absdiff(target, closest) > absdiff(target, currentNode.Value) {
			closest = currentNode.Value
		}
		if target < currentNode.Value {
			currentNode = currentNode.Left

		} else if target > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			break
		}
	}
	return closest
}

func absdiff(a, b int) int {
	out := math.Abs(float64(a) - float64(b))
	return int(out)
}

func main() {
	left := BST{5, nil, nil}
	right := BST{15, nil, nil}
	x := BST{20, &left, &right}

	fmt.Println(x.FindClosestValueOptimize(11))
}
