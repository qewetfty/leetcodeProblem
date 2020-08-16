package greedy

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the
// intervals non-overlapping.
//	Example 1:
//		Input: [[1,2],[2,3],[3,4],[1,3]]
//		Output: 1
//		Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
//	Example 2:
//		Input: [[1,2],[1,2],[1,2]]
//		Output: 2
//		Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.
//	Example 3:
//		Input: [[1,2],[2,3]]
//		Output: 0
//		Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
//	Note:
//		You may assume the interval's end point is always bigger than its start point.
//		Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.

type interval struct {
	intervals [][]int
}

func (this interval) Len() int {
	return len(this.intervals)
}

func (this interval) Less(i, j int) bool {
	starti, startj := this.intervals[i][0], this.intervals[j][0]
	return starti < startj
}

func (this interval) Swap(i, j int) {
	this.intervals[i], this.intervals[j] = this.intervals[j], this.intervals[i]
}

func (this interval) Get(i int) []int {
	return this.intervals[i]
}

// greedy method
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	intervalInts := interval{intervals: intervals}
	sort.Sort(intervalInts)
	prev, count := 0, 0
	for i := 1; i < intervalInts.Len(); i++ {
		prevInt, currInt := intervalInts.Get(prev), intervalInts.Get(i)
		if prevInt[1] > currInt[0] {
			if prevInt[1] > currInt[1] {
				prev = i
			}
			count++
		} else {
			prev = i
		}
	}
	return count
}

// dp method
func eraseOverlapIntervalsDp(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	intervalInts := interval{
		intervals: intervals,
	}
	sort.Sort(intervalInts)
	dp := make([]int, intervalInts.Len())
	dp[0] = 1
	res := 1
	for i := 1; i < intervalInts.Len(); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if !overlap(intervalInts.intervals[j], intervalInts.intervals[i]) {
				max = utils.Max(max, dp[j])
			}
		}
		dp[i] = max + 1
		res = utils.Max(res, dp[i])
	}
	return len(intervals) - res
}

func overlap(i, j []int) bool {
	return i[1] > j[0]
}

func testProblem435() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {1, 3}, {3, 4}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
}
