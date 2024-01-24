package main

import (
	"fmt"
	"strconv"
)

/*
String Compression: Implement a method to perform basic string compression
using the counts of repeated characters. For example, the string
aabcccccaaa would become a2blc5a3, If the "compressed" string would not
become smaller than the original string, your method should return the
original string. You can assume the string has only uppercase and
lowercase letters (a - z).
*/
func compress(s string) string {
	if len(s) <= compressLen(s) {
		return s
	}

	cnt := 0
	var res []rune
	for idx, ch := range s {
		cnt++
		if idx+1 == len(s) || s[idx] != s[idx+1] {
			res = append(res, ch)
			res = append(res, []rune(strconv.Itoa(cnt))...)
			cnt = 0
		}
	}
	return string(res)
} 

func compressLen(s string) int {
	cnt := 0
	res := 0
	for idx := range s {
		cnt++
		if idx+1 == len(s) || s[idx] != s[idx+1] {
			res++
			res += len(strconv.Itoa(cnt))
			cnt = 0
		}
	}
	return res
}

func main() {
	fmt.Println(compress("aabcccccaaa"))
	fmt.Println(compress("abcd"))
}