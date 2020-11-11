package math

import "fmt"

// Given the coordinates of four points in 2D space, return whether the four points could construct a square.
// The coordinate (x,y) of a point is represented by an integer array with two integers.
// 	Example:
//		Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
//		Output: True
//	Note:
//		All the input integers are in the range [-10000, 10000].
//		A valid square has four equal sides with positive length and four equal angles (90-degree angles).
//		Input points have no order.

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return check(p1, p2, p3, p4) || check(p1, p3, p2, p4) || check(p1, p2, p4, p3)
}

func check(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return distance(p1, p2) > 0 && distance(p1, p2) == distance(p2, p3) && distance(p2, p3) == distance(p3, p4) && distance(p3, p4) == distance(p4, p1) && distance(p1, p3) == distance(p2, p4)
}

// 返回两个点距离的平方
func distance(point1, point2 []int) int {
	dx, dy := point1[0]-point2[0], point1[1]-point2[1]
	return dx*dx + dy*dy
}

func testProblem593() {
	fmt.Println(validSquare([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}))
	fmt.Println(validSquare([]int{0, 0}, []int{1, 0}, []int{0, 1}, []int{1, 1}))
	fmt.Println(validSquare([]int{0, 1}, []int{1, 0}, []int{-1, 0}, []int{0, -1}))
	fmt.Println(validSquare([]int{0, 1}, []int{1, 0}, []int{-1, 0}, []int{0, -2}))
	fmt.Println(validSquare([]int{0, 0}, []int{2, 1}, []int{1, -2}, []int{3, -1}))
}
