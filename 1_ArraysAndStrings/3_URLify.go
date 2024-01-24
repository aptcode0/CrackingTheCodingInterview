package main

import (
	"fmt"
	"strings"
)

/*
URLify: Write a method to replace all spaces in a string with '%20'.
*/

func URLify(s string) string {
	s = strings.Trim(s, " ")
	spaces := countSpaces(s)

	res := make([]rune, len(s) + spaces * 2)

	idx := 0
	for _, ch := range s {
		if ch == ' ' {
			res[idx] = '%'
			res[idx+1] = '2'
			res[idx+2] = '0'
			idx += 2
		} else {
			res[idx] = ch
		}
		idx++
	}
	return string(res)
}

func countSpaces(s string) int {
	res := 0
	for _, ch := range s {
		if ch == ' ' {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(URLify("Mr John Smith "))
}