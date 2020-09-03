package main

import (
	"fmt"
	"strings"
)

// 暴力法
func repeatedSubstringPatternFource(s string) bool {
	l := len(s)
	if l <= 1 {
		return false
	}
	for i := 1; i <= l/2; i++ {
		if l%i != 0 {
			continue
		}
		sub := s[:i]
		if strings.Count(s, sub) == l/i {
			return true
		}
	}
	return false
}

// 字符串匹配算法
func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:len(s)*2-1], s)
}

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
	fmt.Println(repeatedSubstringPattern("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"))
}
