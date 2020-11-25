package stack_queue

import "fmt"

// Implement a basic calculator to evaluate a simple expression string.
// The expression string contains only non-negative integers, +, -, *, /
// operators and empty spaces . The integer division should truncate toward zero.
//	Example 1:
//		Input: "3+2*2"
//		Output: 7
//	Example 2:
//		Input: " 3/2 "
//		Output: 1
//	Example 3:
//		Input: " 3+5 / 2 "
//		Output: 5
//	Note:
//		You may assume that the given expression is always valid.
//		Do not use the eval built-in library function.

// 使用数字栈和符号栈
func calculate227(s string) int {
	numberStack, calculatorStack, i, l := make([]int, 0), make([]byte, 0), 0, len(s)
	for i < l {
		curChar := s[i]
		// 如果是数字，就把数字算出来
		if '0' <= curChar && curChar <= '9' {
			curNumber := 0
			for i < l && '0' <= s[i] && s[i] <= '9' {
				curNumber = curNumber*10 + int(s[i]-'0')
				i++
			}
			numberStack = append(numberStack, curNumber)
		} else {
			switch curChar {
			case '+', '-':
				calculatorStack = append(calculatorStack, curChar)
				i++
			case '*':
				// 获取下一个数字并计算结果
				nextNumber := 0
				i++
				// 去除空格
				for s[i] == ' ' {
					i++
				}
				for i < l && '0' <= s[i] && s[i] <= '9' {
					nextNumber = nextNumber*10 + int(s[i]-'0')
					i++
				}
				lastNumber := numberStack[len(numberStack)-1]
				newNumber := lastNumber * nextNumber
				numberStack[len(numberStack)-1] = newNumber
			case '/':
				// 获取下一个数字并计算结果
				nextNumber := 0
				i++
				// 去除空格
				for s[i] == ' ' {
					i++
				}
				for i < l && '0' <= s[i] && s[i] <= '9' {
					nextNumber = nextNumber*10 + int(s[i]-'0')
					i++
				}
				lastNumber := numberStack[len(numberStack)-1]
				newNumber := lastNumber / nextNumber
				numberStack[len(numberStack)-1] = newNumber
			default:
				i++
			}
		}
	}
	// 从前往后加
	for len(calculatorStack) > 0 {
		curCalculator := calculatorStack[0]
		number1, number2 := numberStack[0], numberStack[1]
		if curCalculator == '-' {
			number2 = -number2
		}
		numberStack[1] = number1 + number2
		numberStack, calculatorStack = numberStack[1:], calculatorStack[1:]
	}
	return numberStack[0]
}

func testProblem227() {
	fmt.Println(calculate227("42"))
	fmt.Println(calculate227("1+2-3*4/5+6*7"))
	fmt.Println(calculate227(" 3 / 2 "))
	fmt.Println(calculate227("3+2*2"))
	fmt.Println(calculate227(" 3+5 / 2 "))
}
