package array

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
// You may assume that the intervals were initially sorted according to their start times.
// Example 1:
//		Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
//		Output: [[1,5],[6,9]]
//	Example 2:
//		Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//		Output: [[1,2],[3,10],[12,16]]
//		Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
//	NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

func insert(intervals [][]int, newInterval []int) [][]int {
	i, l, newStart, newEnd := 0, len(intervals), newInterval[0], newInterval[1]
	res := make([][]int, 0)
	// 在newStart之前的，全部加入到结果中
	for i < l && newStart > intervals[i][0] {
		res = append(res, intervals[i])
		i++
	}
	// 判断overlap与最后一个合并规则
	if len(res) == 0 || res[len(res)-1][1] < newStart {
		res = append(res, newInterval)
	} else {
		res[len(res)-1][1] = utils.Max(res[len(res)-1][1], newEnd)
	}
	// 继续向后，如果能够合并则进行合并
	for i < l {
		lastPoint, curPoint := res[len(res)-1], intervals[i]
		// 如果没有overlap了，就直接添加
		if curPoint[0] > lastPoint[1] {
			res = append(res, curPoint)
		} else {
			res[len(res)-1][1] = utils.Max(lastPoint[1], curPoint[1])
		}
		i++
	}
	return res
}

func testProblem57() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
