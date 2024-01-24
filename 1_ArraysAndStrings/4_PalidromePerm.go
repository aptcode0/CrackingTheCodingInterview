package main

import "fmt"

/*
Palindrome Permutation: Given a string, write a function to check if 
it is a permutation of a palin- drome. A palindrome is a word or phrase 
that is the same forwards and backwards. A permutation is a rearrangement 
of letters. 
*/
func isPermPaliUsingMap(s string) bool {
	freqs := buildFreqsMap(s)

	oddSeen := false
	for _, freq := range freqs {
		if freq%2 == 0 {
			continue
		}
		if oddSeen {
			return false
		}
		oddSeen = true
	}
	return true
}

func isPermPaliUsingBitset(s string) bool {
	bitVector := createBitVector(s)
	return bitVector == 0 || isOnlyOneBitSet(bitVector)
}

func createBitVector(s string) int {
	res := 0
	for _, ch := range s {
		pos := (int)(ch - 'a')
		res = toggleBit(res, pos)
	}
	return res
}

func isOnlyOneBitSet(n int) bool {
	return n & (n-1) == 0
}
func toggleBit(n int, pos int) int {
	if isBitSet(n, pos) {
		return unsetBit(n, pos)
	}
	return setBit(n, pos)
}

func isBitSet(n int, pos int) bool {
	return n&(1<<pos) > 0
}

func unsetBit(n int, pos int) int {
	return n & ^(1 << pos)
}

func setBit(n int, pos int) int {
	return n | (1 << pos)
}

func buildFreqsMap(s string) map[rune]int {
	res := map[rune]int{}
	for _, ch := range s {
		res[ch]++
	}
	return res
}


func main() {
	fmt.Println(isPermPaliUsingMap("pattap"))
	fmt.Println(isPermPaliUsingMap("apt"))

	fmt.Println(isPermPaliUsingBitset("pattap"))
	fmt.Println(isPermPaliUsingBitset("apt"))

}