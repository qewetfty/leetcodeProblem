package stack_queue

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Implement a basic calculator to evaluate a simple expression string.
// The expression string may contain open ( and closing parentheses ),
// the plus + or minus sign -, non-negative integers and empty spaces .
// Note:
//	You may assume that the given expression is always valid.
//	Do not use the eval built-in library function.

func calculate(s string) int {
	result, _ := helper(s, 0)
	return result
}

func helper(s string, pos int) (result int, position int) {
	if len(s) == 0 {
		return 0, 0
	}
	stack := new(data.Stack)
	i := pos
	sign := '+'
	if s[i] == '-' {
		sign = '-'
		i = i + 1
	} else if s[i] == '+' {
		i = i + 1
	}
	num := 0
	for i < len(s) {
		c := s[i]
		if isDigit(c) {
			num = 10*num + int(c-'0')
		}
		if c == '(' {
			num, i = helper(s, i+1)
		}
		if !isDigit(c) && c != ' ' || i == len(s)-1 {
			switch sign {
			case '+':
				stack.Push(num)
				break
			case '-':
				stack.Push(-num)
				break
			case '*':
				prev := stack.Top().(int)
				stack.Pop()
				stack.Push(prev * num)
				break
			case '/':
				prev := stack.Top().(int)
				stack.Pop()
				stack.Push(prev / num)
				break
			}
			sign = int32(c)
			num = 0
		}
		if c == ')' {
			break
		}
		i++
	}
	res := 0
	for !stack.IsEmpty() {
		num, _ := stack.Pop()
		res += num.(int)
	}
	return res, i
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func testProblem224() {
	fmt.Println(calculate("2147483647"))
	fmt.Println(calculate("+2 -3 * 8 / 4"))
	fmt.Println(calculate("2 + 3 * 4 - 7"))
	fmt.Println(calculate("-2 * 4 + 6 / 3 + 2"))
	fmt.Println(calculate("2 + (3 +5) * 2"))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
