package dfs_bfs

import "fmt"

// Given an array of integers arr, you are initially positioned at the first index of the array.
// In one step you can jump from index i to index:
//	i + 1 where: i + 1 < arr.length.
//	i - 1 where: i - 1 >= 0.
//	j where: arr[i] == arr[j] and i != j.
//	Return the minimum number of steps to reach the last index of the array.
// Notice that you can not jump outside of the array at any time.
//	Example 1:
//		Input: arr = [100,-23,-23,404,100,23,23,23,3,404]
//		Output: 3
//		Explanation: You need three jumps from index 0 --> 4 --> 3 --> 9. Note that index 9 is the last index of the array.
//	Example 2:
//		Input: arr = [7]
//		Output: 0
//		Explanation: Start index is the last index. You don't need to jump.
//	Example 3:
//		Input: arr = [7,6,9,6,9,6,9,7]
//		Output: 1
//		Explanation: You can jump directly from index 0 to index 7 which is last index of the array.
//	Example 4:
//		Input: arr = [6,1,9]
//		Output: 2
//	Example 5:
//		Input: arr = [11,22,7,7,7,7,7,7,7,22,13]
//		Output: 3
//	Constraints:
//		1 <= arr.length <= 5 * 10^4
//		-10^8 <= arr[i] <= 10^8

func minJumps(arr []int) int {
	l := len(arr)
	if l <= 1 {
		return 0
	}
	// 记录每一个number对应的位置, unique记录数组是否有重复的数据。如果没有，直接返回l-1
	numberMap, unique := make(map[int][]int), true
	// 这里使用倒序遍历主要是应对测试用例[7,7,7,7,7.......7,7,11]
	for i := l - 1; i >= 0; i-- {
		numberMap[arr[i]] = append(numberMap[arr[i]], i)
		if len(numberMap[arr[i]]) >= 2 {
			unique = false
		}
	}
	if unique {
		return l - 1
	}
	queue, step := make([]int, 0), 0
	// 记录哪些位置已经被遍历过
	arrivedMap := make(map[int]bool)
	queue = append(queue, 0)
	arrivedMap[0] = true
	for len(queue) > 0 {
		curLen := len(queue)
		step++
		for i := 0; i < curLen; i++ {
			pos := queue[0]
			queue = queue[1:]
			// 处理下一个元素
			if pos+1 < l {
				if pos+1 == l-1 {
					return step
				} else {
					if !arrivedMap[pos+1] {
						queue = append(queue, pos+1)
						arrivedMap[pos+1] = true
					}
				}
			}
			// 处理前一个元素
			if pos-1 >= 0 {
				if !arrivedMap[pos-1] {
					queue = append(queue, pos-1)
					arrivedMap[pos-1] = true
				}
			}
			// 处理对应位置的
			for _, num := range numberMap[arr[pos]] {
				if num == l-1 {
					return step
				}
				if !arrivedMap[num] {
					queue = append(queue, num)
					arrivedMap[num] = true
				}
			}
		}
	}
	return step
}

func testProblem1345() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
	fmt.Println(minJumps([]int{7}))
	fmt.Println(minJumps([]int{7, 6, 9, 6, 9, 6, 9, 7}))
	fmt.Println(minJumps([]int{6, 1, 9}))
	fmt.Println(minJumps([]int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13}))
}
