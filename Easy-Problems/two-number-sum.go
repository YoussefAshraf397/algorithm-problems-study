package main

import (
	"fmt"
	"sort"
)

// TwoNumberSum1 time: o(n^2) | space: o(1)
func TwoNumberSum1(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++ {
			secondNum := array[j]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

// TwoNumberSum2 time: o(n^2) | space: o(n^2)
func TwoNumberSum2(array []int, target int) []int {
	nums := map[int]bool{}
	for _, num := range array {
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found {
			fmt.Println("found ==>", found)
			fmt.Println("nums in success ==>", nums)
			return []int{potentialMatch, num}
		} else {
			nums[num] = true
			fmt.Println("nums in fails ==>", nums)

		}
	}
	return []int{}
}

// TwoNumberSum3 time: o(n * log(n)) | space: o(1)
func TwoNumberSum3(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}

func main() {
	var result []int
	b := []int{3, 1, 2, 5, 6, 8, 7, 4}
	result = TwoNumberSum3(b, 9)
	fmt.Println(result)
}
