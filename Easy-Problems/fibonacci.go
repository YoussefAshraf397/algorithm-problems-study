package main

import "fmt"

// GetNthFib1 O(2^n) time | O(n) space
func GetNthFib1(n int) int {
	if n == 2 {
		return 1
	} else if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	}
	return GetNthFib1(n-1) + GetNthFib1(n-2)
}

// GetNthFib2 O(n) time | O(n) space
func GetNthFib2(n int) int {
	return helper2(n, map[int]int{
		0: 0,
		1: 1,
		2: 1,
	})
}
func helper2(n int, memoize map[int]int) int {
	if val, found := memoize[n]; found {
		fmt.Println("n in if: ", n)
		fmt.Println("found: ", found)
		fmt.Println("val: ", val)
		return val
	}
	memoize[n] = helper2(n-1, memoize) + helper2(n-2, memoize)
	fmt.Println("n out in if: ", n)
	fmt.Println("memoize[n]: ", memoize[n])
	return memoize[n]
}

// GetNthFib3 O(n) time | O(1) space
func GetNthFib3(n int) int {
	lastTwo := []int{1, 1}
	counter := 3
	for counter <= n {
		nextFib := lastTwo[0] + lastTwo[1]
		lastTwo = []int{lastTwo[1], nextFib}
		counter += 1
	}
	if n == 0 {
		return 0
	}
	if n > 1 {
		return lastTwo[1]
	}
	return lastTwo[0]
}

func main() {
	fmt.Println(GetNthFib3(6))
}
