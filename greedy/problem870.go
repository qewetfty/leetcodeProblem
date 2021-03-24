package greedy

import (
	"fmt"
	"sort"
)

// Given two arrays A and B of equal size, the advantage of A with respect to B is the number of indices i for which A[i] > B[i].
// Return any permutation of A that maximizes its advantage with respect to B.
//	Example 1:
//		Input: A = [2,7,11,15], B = [1,10,4,11]
//		Output: [2,11,7,15]
//	Example 2:
//		Input: A = [12,24,8,32], B = [13,25,32,11]
//		Output: [24,32,8,12]
//	Note:
//		1 <= A.length = B.length <= 10000
//		0 <= A[i] <= 10^9
//		0 <= B[i] <= 10^9

func advantageCount(A []int, B []int) []int {
	l := len(A)
	sortedA, sortedB := make([]int, l), make([]int, l)
	copy(sortedA, A)
	copy(sortedB, B)
	sort.Ints(sortedA)
	sort.Ints(sortedB)
	winMap, loseMap := make(map[int][]int), make([]int, 0)
	positionB := 0
	for _, num := range sortedA {
		if num > sortedB[positionB] {
			winMap[sortedB[positionB]] = append(winMap[sortedB[positionB]], num)
			positionB++
		} else {
			loseMap = append(loseMap, num)
		}
	}

	result := make([]int, 0)
	for _, num := range B {
		if _, ok := winMap[num]; ok {
			winQueue := winMap[num]
			curNum := winQueue[0]
			winQueue = winQueue[1:]
			if len(winQueue) == 0 {
				delete(winMap, num)
			} else {
				winMap[num] = winQueue
			}
			result = append(result, curNum)
		} else {
			curNum := loseMap[0]
			loseMap = loseMap[1:]
			result = append(result, curNum)
		}
	}
	return result
}

func testProblem870() {
	fmt.Println(advantageCount([]int{2, 0, 4, 1, 2}, []int{1, 3, 0, 0, 2}))
	fmt.Println(advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}
