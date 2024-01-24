package main

import "fmt"

/*
Is Unique: Implement an algorithm to determine if a string has all unique characters.
What if you cannot use additional data structures?
*/


func isUniqueUsingMap(s string) bool {
	seen := map[rune]bool{}
	for _, ch := range s {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

func isUniqueUsingBitset(s string) bool {
	seen := 0
	for _, ch := range s {
		mask := 1 << (int)(ch - 'a')
		if seen & mask > 0 {
			return false
		}
		seen |= mask
	}
	return true
}

func main() {
	fmt.Println(isUniqueUsingMap("abcd"))
	fmt.Println(isUniqueUsingMap("abad"))
	fmt.Println(isUniqueUsingBitset("abcd"))
	fmt.Println(isUniqueUsingBitset("abad"))
}