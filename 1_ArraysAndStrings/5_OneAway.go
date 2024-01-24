package main

import "fmt"

/*
One Away: There are three types of edits that can be performed on strings:
insert a character, remove a character, or replace a character.
Given two strings, write a function to check if they are one edit
(or zero edits) away.
*/
func isOneEditAway(s1, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if abs(l1-l2) > 1 {
		return false
	}

	if l1 == l2 {
		return isOneReplace(s1, s2)
	}
	if len(s1) > len(s2) {
		return isOneAdd(s1, s2)
	}
	return isOneAdd(s2, s1)
}

func isOneReplace(s1, s2 string) bool {
	isReplaced := false
	for idx := range s1 {
		if s1[idx] == s2[idx] {
			continue
		}
		if isReplaced {
			return false
		}
		isReplaced = true
	}
	return true
} 

func isOneAdd(s1, s2 string) bool {
	offset := 0
	for idx := range s2 {
		if s1[idx+offset] == s2[idx] {
			continue
		}
		if offset > 0 {
			return false
		}
		offset++
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(isOneEditAway("apt", "at"))
	fmt.Println(isOneEditAway("code", "coode"))
	fmt.Println(isOneEditAway("apt", "app"))
	fmt.Println(isOneEditAway("apt", "craft"))
}