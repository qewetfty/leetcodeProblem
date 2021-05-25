package stack_queue

import (
	"fmt"
	"strconv"
)

// Evaluate the value of an arithmetic expression in Reverse Polish Notation.
// Valid operators are +, -, *, and /. Each operand may be an integer or another expression.
// Note that division between two integers should truncate toward zero.
// It is guaranteed that the given RPN expression is always valid. That means the
// expression would always evaluate to a result, and there will not be any
// division by zero operation.
//	Example 1:
//		Input: tokens = ["2","1","+","3","*"]
//		Output: 9
//		Explanation: ((2 + 1) * 3) = 9
//	Example 2:
//		Input: tokens = ["4","13","5","/","+"]
//		Output: 6
//		Explanation: (4 + (13 / 5)) = 6
//	Example 3:
//		Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
//		Output: 22
//		Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//			= ((10 * (6 / (12 * -11))) + 17) + 5
//			= ((10 * (6 / -132)) + 17) + 5
//			= ((10 * 0) + 17) + 5
//			= (0 + 17) + 5
//			= 17 + 5
//			= 22
//	Constraints:
//		1 <= tokens.length <= 104
//		tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "+":
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			curResult := num1 + num2
			stack = append(stack, curResult)
			break
		case "-":
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			curResult := num1 - num2
			stack = append(stack, curResult)
			break
		case "*":
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			curResult := num1 * num2
			stack = append(stack, curResult)
			break
		case "/":
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			curResult := num1 / num2
			stack = append(stack, curResult)
			break
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func testProblem150() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
