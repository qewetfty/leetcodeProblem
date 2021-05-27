package greedy

import (
	"container/heap"
	"fmt"
	"github.com/leetcodeProblem/data"
	"sort"
)

// There are n different online courses numbered from 1 to n. You are given an
// array courses where courses[i] = [durationi, lastDayi] indicate that the ith
// course should be taken continuously for durationi days and must be finished
// before or on lastDayi.
// You will start on the 1st day and you cannot take two or more courses simultaneously.
// Return the maximum number of courses that you can take.
//	Example 1:
//		Input: courses = [[100,200],[200,1300],[1000,1250],[2000,3200]]
//		Output: 3
//		Explanation:
//			There are totally 4 courses, but you can take 3 courses at most:
//			First, take the 1st course, it costs 100 days so you will finish it on the 100th day, and ready to take the next course on the 101st day.
//			Second, take the 3rd course, it costs 1000 days so you will finish it on the 1100th day, and ready to take the next course on the 1101st day.
//			Third, take the 2nd course, it costs 200 days so you will finish it on the 1300th day.
//			The 4th course cannot be taken now, since you will finish it on the 3300th day, which exceeds the closed date.
//	Example 2:
//		Input: courses = [[1,2]]
//		Output: 1
//	Example 3:
//		Input: courses = [[3,2],[4,3]]
//		Output: 0
//	Constraints:
//		1 <= courses.length <= 104
//		1 <= durationi, lastDayi <= 104

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	priorityQueue := new(data.IntHeapDesc)
	time := 0
	for _, course := range courses {
		if time+course[0] <= course[1] {
			heap.Push(priorityQueue, course[0])
			time += course[0]
		} else {
			if priorityQueue.Len() > 0 {
				top := heap.Pop(priorityQueue).(int)
				if top > course[0] {
					time += course[0] - top
					heap.Push(priorityQueue, course[0])
				} else {
					heap.Push(priorityQueue, top)
				}
			}
		}
	}
	return priorityQueue.Len()
}

func testProblem630() {
	fmt.Println(scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}))
	fmt.Println(scheduleCourse([][]int{{1, 2}}))
	fmt.Println(scheduleCourse([][]int{{3, 2}, {4, 3}}))
}
