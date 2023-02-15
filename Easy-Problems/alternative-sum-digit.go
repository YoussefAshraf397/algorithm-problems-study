package main

import (
	"fmt"
)

func alternateDigitSum(n int) int {
	sign, ans := 1, 0
	for n > 0 {
		ans += (n % 10) * sign
		n = n / 10
		sign *= -1
	}
	return -sign * ans
}

func main() {

	fmt.Println(alternateDigitSum(12))
}
