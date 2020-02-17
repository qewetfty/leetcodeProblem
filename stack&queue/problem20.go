package stack_queue

import (
	"fmt"
	"github.com/leetcodeProblem/data"
	"strings"
)

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
// An input string is valid if:
//	Open brackets must be closed by the same type of brackets.
//	Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.

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

func isValid2(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := new(data.Stack)
	brackets := make(map[string]string, 3)
	brackets[")"] = "("
	brackets["]"] = "["
	brackets["}"] = "{"
	str := []byte(s)
	for _, c := range str {
		if sta, exists := brackets[string(c)]; exists {
			stackChar, err := stack.Pop()
			if err != nil || !strings.EqualFold(sta, stackChar.(string)) {
				return false
			}
		} else {
			stack.Push(string(c))
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
