package greedy

import (
	"fmt"
	"sort"
)

// There are some spherical balloons spread in two-dimensional space. For each
// balloon, provided input is the start and end coordinates of the horizontal
// diameter. Since it's horizontal, y-coordinates don't matter, and hence the
// x-coordinates of start and end of the diameter suffice. The start is always
// smaller than the end. An arrow can be shot up exactly vertically from
// different points along the x-axis. A balloon with xstart and xend bursts by an
// arrow shot at x if xstart ≤ x ≤ xend. There is no limit to the number of
// arrows that can be shot. An arrow once shot keeps traveling up infinitely.
// Given an array points where points[i] = [xstart, xend], return the minimum
// number of arrows that must be shot to burst all balloons.
//	Example 1:
//		Input: points = [[10,16],[2,8],[1,6],[7,12]]
//		Output: 2
//		Explanation: One way is to shoot one arrow for example at x = 6 (bursting the balloons [2,8] and [1,6]) and
//					another arrow at x = 11 (bursting the other two balloons).
//	Example 2:
//		Input: points = [[1,2],[3,4],[5,6],[7,8]]
//		Output: 4
//	Example 3:
//		Input: points = [[1,2],[2,3],[3,4],[4,5]]
//		Output: 2
//	Example 4:
//		Input: points = [[1,2]]
//		Output: 1
//	Example 5:
//		Input: points = [[2,3],[2,3]]
//		Output: 1
//	Constraints:
//		0 <= points.length <= 104
//		points.length == 2
//		-231 <= xstart < xend <= 231 - 1

// 按每一个结尾排序
func findMinArrowShots(points [][]int) int {
	l := len(points)
	if l < 2 {
		return l
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	res, end := 1, points[0][1]
	for i := 1; i < l; i++ {
		if points[i][0] > end {
			res, end = res+1, points[i][1]
		}
	}
	return res
}

// 按第一个排序，击败37.5%
func findMinArrowShots2(points [][]int) int {
	l := len(points)
	if l == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	res, i, right := 0, 1, points[0][1]
	for i < l {
		if points[i][1] < right {
			right = points[i][1]
		}
		if points[i][0] > right {
			res, right = res+1, points[i][1]
		}
		i++
	}
	return res + 1
}

func testProblem452() {
	fmt.Println(findMinArrowShots([][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}))
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}}))
	fmt.Println(findMinArrowShots([][]int{{2, 3}, {2, 3}}))
}
