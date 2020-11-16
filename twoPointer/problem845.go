package twoPointer

import "fmt"

// Let's call any (contiguous) subarray B (of A) a mountain if the following properties hold:
//	B.length >= 3
//	There exists some 0 < i < B.length - 1 such that B[0] < B[1] < ... B[i-1] < B[i] > B[i+1] > ... > B[B.length - 1]
//	(Note that B could be any subarray of A, including the entire array A.)
// Given an array A of integers, return the length of the longest mountain.
// Return 0 if there is no mountain.
//	Example 1:
//		Input: [2,1,4,7,3,2,5]
//		Output: 5
//		Explanation: The largest mountain is [1,4,7,3,2] which has length 5.
//	Example 2:
//		Input: [2,2,2]
//		Output: 0
//		Explanation: There is no mountain.
//	Note:
//		0 <= A.length <= 10000
//		0 <= A[i] <= 10000
//	Follow up:
//		Can you solve it using only one pass?
//		Can you solve it in O(1) space?

// two pointer method
func longestMountain(A []int) int {
	l, left, res := len(A), 0, 0
	for left+2 < l {
		right := left + 1
		if A[left] < A[left+1] {
			for right+1 < l && A[right] < A[right+1] {
				right++
			}
			if right+1 < l && A[right] > A[right+1] {
				for right+1 < l && A[right] > A[right+1] {
					right++
				}
				if right-left+1 > res {
					res = right - left + 1
				}
			} else {
				right++
			}
		}
		left = right
	}
	return res
}

// dp method
func longestMountain2(A []int) int {
	l := len(A)
	// 枚举左侧山脉延伸长度
	left := make([]int, l)
	for i := 1; i < l; i++ {
		if A[i-1] < A[i] {
			left[i] = left[i-1] + 1
		}
	}
	// 枚举右侧山脉延伸长度
	right := make([]int, l)
	for i := l - 2; i >= 0; i-- {
		if A[i+1] < A[i] {
			right[i] = right[i+1] + 1
		}
	}
	// 枚举最长山脉
	res := 0
	for i, l := range left {
		r := right[i]
		for r > 0 && l > 0 && r+l+1 > res {
			res = l + r + 1
		}
	}
	return res
}

func testProblem845() {
	fmt.Println(longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
	fmt.Println(longestMountain([]int{2, 2, 2}))
}
