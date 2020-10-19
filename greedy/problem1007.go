package greedy

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// In a row of dominoes, A[i] and B[i] represent the top and bottom halves of the
// ith domino. (A domino is a tile with two numbers from 1 to 6 - one on each
// half of the tile.)
// We may rotate the ith domino, so that A[i] and B[i] swap values.
// Return the minimum number of rotations so that all the values in A are the same, or all the values in B are the same.
// If it cannot be done, return -1.
//	Example 1:
//		Input: A = [2,1,2,4,2,2], B = [5,2,6,2,3,2]
//		Output: 2
//		Explanation:
//			The first figure represents the dominoes as given by A and B: before we do any
//			rotations. If we rotate the second and fourth dominoes, we can make every
//			value in the top row equal to 2, as indicated by the second figure.
//	Example 2:
//		Input: A = [3,5,1,2,3], B = [3,6,3,3,4]
//		Output: -1
//		Explanation:
//			In this case, it is not possible to rotate the dominoes to make one row of values equal.
//	Constraints:
//		2 <= A.length == B.length <= 2 * 104
//		1 <= A[i], B[i] <= 6

// 如果能旋转，只需判断第一个两个元素是否能够旋转，不能旋转则一定是整个都不能旋转的，返回-1即可
func minDominoRotations(A []int, B []int) int {
	l := len(A)
	r := check(A[0], l, A, B)
	if r != -1 || A[0] == B[0] {
		return r
	}
	return check(B[0], l, A, B)
}

func check(x, l int, A, B []int) int {
	rotationsA, rotationsB := 0, 0
	for i := 0; i < l; i++ {
		if A[i] != x && B[i] != x {
			return -1
		} else if A[i] != x {
			rotationsA++
		} else if B[i] != x {
			rotationsB++
		}
	}
	return utils.Min(rotationsA, rotationsB)
}

// 二维数组保存每一个点数出现的位置，然后进行判断，击败52%，空间复杂度较高。
func minDominoRotations2(A []int, B []int) int {
	l, a1, b1 := len(A), make([][]int, 6), make([][]int, 6)
	// 创建二维数组保存每个点数的位置
	for i := 0; i < 6; i++ {
		a1[i], b1[i] = make([]int, 0), make([]int, 0)
	}
	for i := 0; i < l; i++ {
		a1[A[i]-1], b1[B[i]-1] = append(a1[A[i]-1], i), append(b1[B[i]-1], i)
	}
	res := math.MaxInt32
	for i := 0; i < 6; i++ {
		a, b := a1[i], b1[i]
		// 每个点数的总和小于长度l的话，跳过
		if len(a)+len(b) < l {
			continue
		}
		// a是长度长的、
		if len(a) < len(b) {
			a, b = b, a
		}
		ia, ib, move, find := 0, 0, 0, false
		// 判断从0到l是否每一个都有
		for j := 0; j < l; j++ {
			findB := false
			find = false
			// 如果短的找到了，可能会交换，将findB变为true
			if ib < len(b) && b[ib] == j {
				find, findB, ib = true, true, ib+1
			}
			// 如果a存在的话，则不需要交换a,b，将findB变为false
			if ia < len(a) && a[ia] == j {
				find, findB, ia = true, false, ia+1
			}
			// 如果没找到，直接跳出
			if !find {
				res = math.MaxInt32
				break
			}
			if findB {
				move++
			}
		}
		if find {
			res = utils.Min(res, move)
		}
	}
	if res == math.MaxInt32 {
		return -1
	} else {
		return res
	}
}

func testProblem1007() {
	fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}))
	fmt.Println(minDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}))
	fmt.Println(minDominoRotations([]int{1, 1, 1, 1, 1}, []int{3, 6, 3, 3, 4}))
}
