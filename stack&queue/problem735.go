package stack_queue

import "fmt"

// We are given an array asteroids of integers representing asteroids in a row.
// For each asteroid, the absolute value represents its size, and the sign
// represents its direction (positive meaning right, negative meaning left). Each
// asteroid moves at the same speed.
// Find out the state of the asteroids after all collisions. If two asteroids
// meet, the smaller one will explode. If both are the same size, both will
// explode. Two asteroids moving in the same direction will never meet.
//	Example 1:
//		Input: asteroids = [5,10,-5]
//		Output: [5,10]
//		Explanation: The 10 and -5 collide resulting in 10.  The 5 and 10 never collide.
//	Example 2:
//		Input: asteroids = [8,-8]
//		Output: []
//		Explanation: The 8 and -8 collide exploding each other.
//	Example 3:
//		Input: asteroids = [10,2,-5]
//		Output: [10]
//		Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.
//	Example 4:
//		Input: asteroids = [-2,-1,1,2]
//		Output: [-2,-1,1,2]
//		Explanation: The -2 and -1 are moving left, while the 1 and 2 are moving right. Asteroids moving the same
//			direction never meet, so no asteroids will meet each other.
//	Constraints:
//		1 <= asteroids <= 104
//		-1000 <= asteroids[i] <= 1000
//		asteroids[i] != 0

func asteroidCollision(asteroids []int) []int {
	l, stack := len(asteroids), make([]int, 0)
	for i := 0; i < l; i++ {
		current := asteroids[i]
		// 与栈顶元素进行对比，共四种情况：
		//	1. 正、正：直接入栈
		//	2. 正、负：出栈，留下绝对值大的或者都湮灭
		//	3. 负、正：直接入栈
		//	4. 负、负：直接入栈
		// 判断碰撞结果
		for len(stack) > 0 && stack[len(stack)-1] > 0 && current < 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if -current == top {
				current = 0
			} else if -current < top {
				current = top
			}
		}
		if current != 0 {
			stack = append(stack, current)
		}
	}
	return stack
}

func testProblem735() {
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -2}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}))
}
