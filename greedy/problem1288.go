package greedy

import (
	"fmt"
	"sort"
)

// Given a list of intervals, remove all intervals that are covered by another interval in the list.
// Interval [a,b) is covered by interval [c,d) if and only if c <= a and b <= d.
// After doing so, return the number of remaining intervals.
//	Example 1:
//		Input: intervals = [[1,4],[3,6],[2,8]]
//		Output: 2
//		Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
//	Example 2:
//		Input: intervals = [[1,4],[2,3]]
//		Output: 1
//	Example 3:
//		Input: intervals = [[0,10],[5,12]]
//		Output: 2
//	Example 4:
//		Input: intervals = [[3,10],[4,10],[5,11]]
//		Output: 2
//	Example 5:
//		Input: intervals = [[1,2],[1,4],[3,4]]
//		Output: 1
//	Constraints:
//		1 <= intervals.length <= 1000
//		intervals[i].length == 2
//		0 <= intervals[i][0] < intervals[i][1] <= 10^5
//		All the intervals are unique.

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})
	count, prevEnd := 0, 0
	for _, interval := range intervals {
		if interval[1] > prevEnd {
			count, prevEnd = count+1, interval[1]
		}
	}
	return count
}

func testProblem1288() {
	fmt.Println(removeCoveredIntervals([][]int{{66672, 75156}, {59890, 65654}, {92950, 95965}, {9103, 31953},
		{54869, 69855}, {33272, 92693}, {52631, 65356}, {43332, 89722}, {4218, 57729}, {20993, 92876}}))
	fmt.Println(removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}}))
	fmt.Println(removeCoveredIntervals([][]int{{1, 4}, {2, 3}}))
	fmt.Println(removeCoveredIntervals([][]int{{0, 10}, {5, 12}}))
	fmt.Println(removeCoveredIntervals([][]int{{3, 10}, {4, 10}, {5, 11}}))
	fmt.Println(removeCoveredIntervals([][]int{{1, 2}, {1, 4}, {3, 4}}))
}
