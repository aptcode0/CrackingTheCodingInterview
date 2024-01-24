package main

import "fmt"

/*
Check Permutation: Given two strings, write a method to decide if 
one is a permutation of the other.
*/

func isPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	freqs := buildFreqsMap(s1)

	for _, ch := range s2 {
		freqs[ch]--
		if freqs[ch] < 0 {
			return false
		}
	}
	return true
}

func buildFreqsMap(s string) map[rune]int {
	res := map[rune]int{}
	for _, ch := range s {
		res[ch]++
	}
	return res
}

func main() {
	fmt.Println(isPermutation("abc", "bca"))
	fmt.Println(isPermutation("abc", "dca"))
}