package main

import (
	"fmt"
)

func reverseArray(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseArray(input[1:]), input[0])
}

func addToArrayForm(num []int, k int) []int {
	var result []int

	for i := len(num) - 1; i >= 0; i -= 1 {
		result = append(result, (num[i]+k)%10)
		k = (num[i] + k) / 10
	}
	for k > 0 {
		result = append(result, k%10)
		k /= 10
	}
	return reverseArray(result)
}

/*
TRACE:
num = 126
k = 516

1- 6+516 = 522 % 10 = 2
res => {2}
k=522/10 = 52

2- 2 + 52 = 54 % 10 = 4
res => {2,4}
k = 54 % 10 = 5

3- 1 + 5 = 6 % 10 = 6
res => {2,4,6}
k = 0
*/

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 6}, 516))
}
