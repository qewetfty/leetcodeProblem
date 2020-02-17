package stack_queue

import (
	"fmt"
	"github.com/leetcodeProblem/data"
	"math"
)

// Given n non-negative integers representing an elevation map where the width of each bar is 1,
// compute how much water it is able to trap after raining.
// The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
// In this case, 6 units of rain water (blue section) are being trapped.
//	Example:
//		Input: [0,1,0,2,1,0,1,3,2,1,2,1]
//		Output: 6

func trap(height []int) int {
	wholeArea := 0
	curr := 0
	stack := new(data.Stack)
	for curr < len(height) {
		for !stack.IsEmpty() && height[curr] > height[stack.Top().(int)] {
			top := stack.Top().(int)
			stack.Pop()
			if stack.IsEmpty() {
				break
			}
			h := int(math.Min(float64(height[curr]), float64(height[stack.Top().(int)]))) - height[top]
			w := curr - stack.Top().(int) - 1
			wholeArea += h * w
		}
		stack.Push(curr)
		curr++
	}
	return wholeArea
}

func testProblem42() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(a))
	b := []int{5, 4, 1, 2}
	fmt.Println(trap(b))
	c := []int{9, 6, 8, 8, 5, 6, 3}
	fmt.Println(trap(c))
}
