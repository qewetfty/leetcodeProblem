package stack_queue

import (
	"fmt"
	"strconv"
)

// Given a balanced parentheses string S, compute the score of the string based on the following rule:
//	() has score 1
//	AB has score A + B, where A and B are balanced parentheses strings.
//	(A) has score 2 * A, where A is a balanced parentheses string.
//	Example 1:
//		Input: "()"
//		Output: 1
//	Example 2:
//		Input: "(())"
//		Output: 2
//	Example 3:
//		Input: "()()"
//		Output: 2
//	Example 4:
//		Input: "(()(()))"
//		Output: 6
//	Note:
//		S is a balanced parentheses string, containing only ( and ).
//		2 <= S.length <= 50

func scoreOfParentheses(S string) int {
	stack, l := make([]string, 0), len(S)
	for i := 0; i < l; i++ {
		char := S[i]
		if char == ')' {
			var score int
			if stack[len(stack)-1] == "(" {
				score = 1
				stack = stack[:len(stack)-1]
			} else {
				j := len(stack) - 1
				for j >= 0 && stack[j] != "(" {
					curScore, _ := strconv.Atoi(stack[j])
					score += curScore
					j--
				}
				score *= 2
				stack = stack[:j]
			}
			stack = append(stack, strconv.Itoa(score))

		} else {
			stack = append(stack, string(char))
		}
	}
	result := 0
	for _, s := range stack {
		score, _ := strconv.Atoi(s)
		result += score
	}
	return result
}

func testProblem856() {
	fmt.Println(scoreOfParentheses("(()(()))"))
	fmt.Println(scoreOfParentheses("()"))
	fmt.Println(scoreOfParentheses("(())"))
	fmt.Println(scoreOfParentheses("()()"))
}
