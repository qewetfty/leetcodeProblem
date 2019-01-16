package stack

import (
	"errors"
	"fmt"
	"strings"
)

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	ss := []byte(s)
	stack := new(Stack)
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

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s *Stack) Push(val interface{}) {
	*s = append(*s, val)
}

func (s Stack) Top() interface{} {
	if s.Len() == 0 {
		return nil
	}
	return s[s.Len()-1]
}

func (s *Stack) Pop() (interface{}, error) {
	stack := *s
	if stack.Len() == 0 {
		return nil, errors.New("Stack is empty.")
	}
	val := stack[stack.Len()-1]
	*s = stack[:stack.Len()-1]
	return val, nil
}

func (s Stack) IsEmpty() bool {
	return s.Len() == 0
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
