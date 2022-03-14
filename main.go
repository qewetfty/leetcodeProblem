package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := make([]string, 0)
	splitPath := strings.Split(path, "/")
	for _, p := range splitPath {
		if ".." == p {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if "." == p || "" == p {
			continue
		} else {
			stack = append(stack, p)
		}
	}
	result := "/" + strings.Join(stack, "/")
	return result
}

func main() {
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/home///foo/"))
}
