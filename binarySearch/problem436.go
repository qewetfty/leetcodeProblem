package binarySearch

import (
	"fmt"
	"sort"
)

// Given a set of intervals, for each of the interval i, check if there exists an interval j whose start point is bigger
// than or equal to the end point of the interval i, which can be called that j is on the "right" of i.
// For any interval i, you need to store the minimum interval j's index, which means that the interval j has the minimum
// start point to build the "right" relationship for interval i. If the interval j doesn't exist, store -1 for the
// interval i. Finally, you need output the stored value of each interval as an array.
//	Note:
//		You may assume the interval's end point is always bigger than its start point.
//		You may assume none of these intervals have the same start point.
//	Example 1:
//		Input: [ [1,2] ]
//		Output: [-1]
//		Explanation: There is only one interval in the collection, so it outputs -1.
//	Example 2:
//		Input: [ [3,4], [2,3], [1,2] ]
//		Output: [-1, 0, 1]
//		Explanation: There is no satisfied "right" interval for [3,4].
//					For [2,3], the interval [3,4] has minimum-"right" start point;
//					For [1,2], the interval [2,3] has minimum-"right" start point.
//	Example 3:
//		Input: [ [1,4], [2,3], [3,4] ]
//		Output: [-1, 2, -1]
//		Explanation: There is no satisfied "right" interval for [1,4] and [3,4].
//					For [2,3], the interval [3,4] has minimum-"right" start point.
//		NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

func findRightInterval(intervals [][]int) []int {
	l := len(intervals)
	if l == 0 {
		return []int{}
	}
	if l == 1 {
		return []int{-1}
	}
	res := make([]int, 0)
	newInterval := new(interval)
	newInterval.nums, newInterval.indexes = make([][]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		newInterval.nums[i] = []int{intervals[i][0], intervals[i][1]}
		newInterval.indexes[i] = i
	}
	sort.Sort(newInterval)
	for _, interval := range intervals {
		right := interval[1]
		lo, hi := 0, l
		for lo < hi {
			mid := (lo + hi) / 2
			if newInterval.nums[mid][0] >= right {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		if hi == l {
			res = append(res, -1)
		} else {
			res = append(res, newInterval.indexes[hi])
		}
	}
	return res
}

type interval struct {
	nums    [][]int
	indexes []int
}

func (in *interval) Less(i, j int) bool {
	return in.nums[i][0] < in.nums[j][0]
}

func (in *interval) Swap(i, j int) {
	in.nums[i], in.nums[j] = in.nums[j], in.nums[i]
	in.indexes[i], in.indexes[j] = in.indexes[j], in.indexes[i]
}

func (in *interval) Len() int {
	return len(in.nums)
}

func testProblem436() {
	fmt.Println(findRightInterval([][]int{{1, 2}}))
	fmt.Println(findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}}))
	fmt.Println(findRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}}))
}
