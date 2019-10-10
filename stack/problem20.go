package stack

import (
	"fmt"
	"github.com/leetcodeProblem/data"
	"strings"
)

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	ss := []byte(s)
	stack := new(data.Stack)
	for i := 0; i < len(s); i++ {
		nextString := string(ss[i])
		match := false
		switch nextString {
		case ")":
			for stack.Len() > 0 {
				x, _ := stack.Pop()
				other := "{["
				if strings.Contains(other, x.(string)) {
					return false
				}
				if strings.EqualFold(x.(string), "(") {
					match = true
					break
				}
			}
			if !match {
				return false
			}
		case "]":
			for stack.Len() > 0 {
				x, _ := stack.Pop()
				other := "({"
				if strings.Contains(other, x.(string)) {
					return false
				}
				if strings.EqualFold(x.(string), "[") {
					match = true
					break
				}
			}
			if !match {
				return false
			}
		case "}":
			for stack.Len() > 0 {
				x, _ := stack.Pop()
				other := "(["
				if strings.Contains(other, x.(string)) {
					return false
				}
				if strings.EqualFold(x.(string), "{") {
					match = true
					break
				}
			}
			if !match {
				return false
			}
		default:
			stack.Push(nextString)
		}
	}
	return stack.IsEmpty()
}

func testProblem20() {
	input1 := "()"
	fmt.Println(isValid(input1))
	input2 := "()[]{}"
	fmt.Println(isValid(input2))
	input3 := "({"
	fmt.Println(isValid(input3))
	input4 := "([)]"
	fmt.Println(isValid(input4))
	input5 := "{[]}"
	fmt.Println(isValid(input5))
	input6 := "([)"
	fmt.Println(isValid(input6))
}
