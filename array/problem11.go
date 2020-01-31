package array

import (
	"fmt"
	"math"
)

// Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
// n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
// Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//
//Note: You may not slant the container and n is at least 2.

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	sumArea := 0
	for start < end {
		left := height[start]
		right := height[end]
		containerHeight := int(math.Min(float64(left), float64(right)))
		containerLength := end - start
		area := containerHeight * containerLength
		if area > sumArea {
			sumArea = area
		}
		if left < right {
			start++
		} else {
			end--
		}
	}
	return sumArea
}

func testProblem11() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(a))
}
