package main

import "fmt"

// IsValidSubsequence O(n) time | O(1) space
func IsValidSubsequence(array []int, sequence []int) bool {
	arrIdx := 0
	seqIdx := 0
	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx += 1
		}
		arrIdx += 1
	}
	return seqIdx == len(sequence)
}

func isSubsequence(s string, t string) bool {
	arrIdx := 0
	seqIdx := 0
	for arrIdx < len(t) && seqIdx < len(s) {
		if t[arrIdx] == s[seqIdx] {
			seqIdx += 1
		}
		arrIdx += 1
	}
	return seqIdx == len(s)
}

func main() {
	//arr := []int{3, 1, 2, 5, 6, 8, 7, 4}
	//seq := []int{5, 8}

	fmt.Println(isSubsequence("axc", "ahbgdc"))

	//fmt.Println(IsValidSubsequence(arr, seq))
}
