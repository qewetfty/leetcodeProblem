package array

import (
	"fmt"
	"sort"
)

// Given an array of intervals where intervals[i] = [starti, endi], merge all
// overlapping intervals, and return an array of the non-overlapping intervals
// that cover all the intervals in the input.
//	Example 1:
//		Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
//		Output: [[1,6],[8,10],[15,18]]
//		Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
//	Example 2:
//		Input: intervals = [[1,4],[4,5]]
//		Output: [[1,5]]
//		Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//	Constraints:
//		1 <= intervals.length <= 104
//		intervals[i].length == 2
//		0 <= starti <= endi <= 104

func merge56(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[j][0] < intervals[i][0]
	})
	l, start, end := len(intervals), intervals[0][0], intervals[0][1]
	res := make([][]int, 0)
	for i := 1; i < l; i++ {
		if intervals[i][0] > end {
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		}
	}
	res = append(res, []int{start, end})
	return res
}

func testProblem56() {
	fmt.Println(merge56([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge56([][]int{{1, 4}, {4, 5}}))
}
