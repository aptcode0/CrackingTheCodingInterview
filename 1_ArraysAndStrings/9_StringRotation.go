package main

import (
	"fmt"
	"strings"
)

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