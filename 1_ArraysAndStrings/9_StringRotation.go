package main

import (
	"fmt"
	"strings"
)
/*
String Rotation: Given two strings, s1 and s2, write code to check if s2 is a rotation of s1
using only one call to isSubstring [e.g., "waterbottle" is a rotation oP'erbottlewat"),
*/
func isRotated(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1s1 := s1 + s1 

	return strings.Index(s1s1, s2) >= 0
}

func main() {
	fmt.Println(isRotated("waterbottle", "erbottlewat"))
	fmt.Println(isRotated("water", "retaw"))
}