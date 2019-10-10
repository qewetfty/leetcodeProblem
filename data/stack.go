package data

import "errors"

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
