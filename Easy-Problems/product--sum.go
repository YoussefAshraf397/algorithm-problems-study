package main

import "fmt"

type SpecialArray []interface{}

// ProductSum O(n) time | O(d) space - where n is the total number of elements in
// including sub-elements, and d is the greatest depth of "special" ar
func ProductSum(array SpecialArray) int {
	return helper(array, 1)
}

func helper(array SpecialArray, multiplier int) int {
	sum := 0
	for _, el := range array {
		if cast, ok := el.(SpecialArray); ok {
			fmt.Println("cast 1 : ", cast)
			sum += helper(cast, multiplier+1)
		} else if cast, ok := el.(int); ok {
			fmt.Println("cast 2 : ", cast)

			sum += cast
		}
	}
	return sum * multiplier
}

func main() {
	arr := SpecialArray{5, 2, SpecialArray{7, -1}, 3, SpecialArray{6, SpecialArray{-13, 8}, 4}}
	fmt.Println(ProductSum(arr))
}
