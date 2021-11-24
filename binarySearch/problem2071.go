package binarySearch

import (
	"github.com/leetcodeProblem/utils"
	"sort"
)

//You have n tasks and m workers. Each task has a strength requirement stored
//in a 0-indexed integer array tasks, with the iᵗʰ task requiring tasks[i] strength
//to complete. The strength of each worker is stored in a 0-indexed integer array
//workers, with the jᵗʰ worker having workers[j] strength. Each worker can only
//be assigned to a single task and must have a strength greater than or equal to
//the task's strength requirement (i.e., workers[j] >= tasks[i]).
//
// Additionally, you have pills magical pills that will increase a worker's
//strength by strength. You can decide which workers receive the magical pills,
//however, you may only give each worker at most one magical pill.
//
// Given the 0-indexed integer arrays tasks and workers and the integers pills
//and strength, return the maximum number of tasks that can be completed.
//
//
// Example 1:
//
//
//Input: tasks = [3,2,1], workers = [0,3,3], pills = 1, strength = 1
//Output: 3
//Explanation:
//We can assign the magical pill and tasks as follows:
//- Give the magical pill to worker 0.
//- Assign worker 0 to task 2 (0 + 1 >= 1)
//- Assign worker 1 to task 1 (3 >= 2)
//- Assign worker 2 to task 0 (3 >= 3)
//
//
// Example 2:
//
//
//Input: tasks = [5,4], workers = [0,0,0], pills = 1, strength = 5
//Output: 1
//Explanation:
//We can assign the magical pill and tasks as follows:
//- Give the magical pill to worker 0.
//- Assign worker 0 to task 0 (0 + 5 >= 5)
//
//
// Example 3:
//
//
//Input: tasks = [10,15,30], workers = [0,10,10,10,10], pills = 3, strength = 10
//
//Output: 2
//Explanation:
//We can assign the magical pills and tasks as follows:
//- Give the magical pill to worker 0 and worker 1.
//- Assign worker 0 to task 0 (0 + 10 >= 10)
//- Assign worker 1 to task 1 (10 + 10 >= 15)
//
//
// Example 4:
//
//
//Input: tasks = [5,9,8,5,9], workers = [1,6,4,2,6], pills = 1, strength = 5
//Output: 3
//Explanation:
//We can assign the magical pill and tasks as follows:
//- Give the magical pill to worker 2.
//- Assign worker 1 to task 0 (6 >= 5)
//- Assign worker 2 to task 2 (4 + 5 >= 8)
//- Assign worker 4 to task 3 (6 >= 5)
//
//
//
// Constraints:
//
//
// n == tasks.length
// m == workers.length
// 1 <= n, m <= 5 * 10⁴
// 0 <= pills <= m
// 0 <= tasks[i], workers[j], strength <= 10⁹
//
// Related Topics Array Binary Search Greedy Queue Sorting Monotonic Queue 👍 11
//1 👎 7

//leetcode submit region begin(Prohibit modification and deletion)
func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n, m := len(tasks), len(workers)
	sort.Ints(tasks)
	sort.Ints(workers)

	check := func(mid int) bool {
		p := pills
		deque := make([]int, 0)
		ptr := m - 1
		for i := mid - 1; i >= 0; i-- {
			for ptr >= m-mid && workers[ptr]+strength >= tasks[i] {
				deque = append(deque, workers[ptr])
				ptr--
			}
			if len(deque) == 0 {
				return false
			} else if deque[0] >= tasks[i] {
				deque = deque[1:]
			} else {
				if p == 0 {
					return false
				}
				p--
				deque = deque[:len(deque)-1]
			}
		}
		return true
	}

	lo, hi := 0, utils.Min(m, n)
	result := 0
	for lo <= hi {
		mid := (lo + hi) >> 1
		if check(mid) {
			result = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}
