package dfs_bfs

import "fmt"

// Given an array of non-negative integers arr, you are initially positioned at
// start index of the array. When you are at index i, you can jump to i + arr[i]
// or i - arr[i], check if you can reach to any index with value 0.
// Notice that you can not jump outside of the array at any time.
//	Example 1:
//		Input: arr = [4,2,3,0,3,1,2], start = 5
//		Output: true
//		Explanation:
//			All possible ways to reach at index 3 with value 0 are:
//			index 5 -> index 4 -> index 1 -> index 3
//			index 5 -> index 6 -> index 4 -> index 1 -> index 3
//	Example 2:
//		Input: arr = [4,2,3,0,3,1,2], start = 0
//		Output: true
//		Explanation:
//			One possible way to reach at index 3 with value 0 is:
//			index 0 -> index 4 -> index 1 -> index 3
//	Example 3:
//		Input: arr = [3,0,2,1,2], start = 2
//		Output: false
//		Explanation: There is no way to reach at index 1 with value 0.
//	Constraints:
//		1 <= arr.length <= 5 * 10^4
//		0 <= arr[i] < arr.length
//		0 <= start < arr.length

func canReach(arr []int, start int) bool {
	l := len(arr)
	queue, visits := make([]int, 0), make([]bool, l)
	queue, visits[start] = append(queue, start), true
	for len(queue) > 0 {
		pos := queue[0]
		value := arr[pos]
		queue = queue[1:]
		if value == 0 {
			return true
		}
		if pos+value < l && !visits[pos+value] {
			queue, visits[pos+value] = append(queue, pos+value), true
		}
		if pos-value >= 0 && !visits[pos-value] {
			queue, visits[pos-value] = append(queue, pos-value), true
		}
	}
	return false
}

func testProblem1306() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0))
	fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))
}
