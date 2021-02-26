package stack_queue

import "fmt"

// Given two sequences pushed and popped with distinct values, return true if and
// only if this could have been the result of a sequence of push and pop
// operations on an initially empty stack.
//	Example 1:
//		Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
//		Output: true
//		Explanation: We might do the following sequence:
//			push(1), push(2), push(3), push(4), pop() -> 4,
//			push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//	Example 2:
//		Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
//		Output: false
//		Explanation: 1 cannot be popped before 2.
//	Constraints:
//		0 <= pushed.length == popped.length <= 1000
//		0 <= pushed[i], popped[i] < 1000
//		pushed is a permutation of popped.
//		pushed and popped have distinct values.

func validateStackSequences(pushed []int, popped []int) bool {
	l := len(pushed)
	i, j := 0, 0
	stack := make([]int, 0)
	// 处理push进去的，以及中间可能pop出来的
	for i < l {
		stackLen := len(stack)
		for stackLen > 0 && j < l && stack[stackLen-1] == popped[j] {
			j++
			stackLen--
		}
		stack = stack[:stackLen]
		stack = append(stack, pushed[i])
		i++
	}
	// push完成之后处理剩下需要pop的
	k := len(stack)
	for j < l {
		if stack[k-1] != popped[j] {
			return false
		}
		k--
		j++
	}
	stack = stack[:k]
	return len(stack) == 0
}

func testProblem946() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
