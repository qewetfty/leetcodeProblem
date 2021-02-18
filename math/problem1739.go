package math

import "fmt"

// You have a cubic storeroom where the width, length, and height of the room are
// all equal to n units. You are asked to place n boxes in this room where each
// box is a cube of unit side length. There are however some rules to placing the
// boxes:
//	You can place the boxes anywhere on the floor.
//	If box x is placed on top of the box y, then each side of the four vertical sides of the box y must either be adjacent to another box or to a wall.
// Given an integer n, return the minimum possible number of boxes touching the floor.
//	Example 1:
//		Input: n = 3
//		Output: 3
//		Explanation: The figure above is for the placement of the three boxes.
//		These boxes are placed in the corner of the room, where the corner is on the left side.
//	Example 2:
//		Input: n = 4
//		Output: 3
//		Explanation: The figure above is for the placement of the four boxes.
//		These boxes are placed in the corner of the room, where the corner is on the left side.
//	Example 3:
//		Input: n = 10
//		Output: 6
//		Explanation: The figure above is for the placement of the ten boxes.
//		These boxes are placed in the corner of the room, where the corner is on the back side.
//	Constraints:
//		1 <= n <= 109

// O(n^1/2)解法。可改为使用二分法求level和i
func minimumBoxes(n int) int {
	level := 1
	for level*(level+1)*(level+2)/6 <= n {
		level++
	}
	level--
	remain := n - level*(level+1)*(level+2)/6
	i := 0
	for remain > 0 {
		remain -= i
		i++
	}
	if i > 0 {
		i--
	}
	return level*(level+1)/2 + i
}

func testProblem1739() {
	fmt.Println(minimumBoxes(3))
	fmt.Println(minimumBoxes(4))
	fmt.Println(minimumBoxes(10))
	fmt.Println(minimumBoxes(100000000))
}
