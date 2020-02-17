package stack_queue

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Given n non-negative integers representing the histogram's bar height where the width of each bar is 1,
// find the area of largest rectangle in the histogram.

func largestRectangleArea(heights []int) int {
	stack := new(data.Stack)
	stack.Push(-1)
	maxArea := 0
	for i, num := range heights {
		for stack.Top().(int) != -1 && heights[stack.Top().(int)] >= num {
			val, _ := stack.Pop()
			area := heights[val.(int)] * (i - stack.Top().(int) - 1)
			if area > maxArea {
				maxArea = area
			}
		}
		stack.Push(i)
	}
	for stack.Top().(int) != -1 {
		val, _ := stack.Pop()
		area := heights[val.(int)] * (len(heights) - stack.Top().(int) - 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func testProblem84() {
	a := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(a))
	b := []int{6, 7, 5, 2, 4, 5, 9, 3}
	fmt.Println(largestRectangleArea(b))
}
