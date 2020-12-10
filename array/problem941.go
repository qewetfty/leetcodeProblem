package array

import "fmt"

// Given an array of integers arr, return true if and only if it is a valid mountain array.
// Recall that arr is a mountain array if and only if:
//		arr.length >= 3
//		There exists some i with 0 < i < arr.length - 1 such that:
//		arr[0] < arr[1] < ... < arr[i - 1] < A[i]
//		arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
//	Example 1:
//		Input: arr = [2,1]
//		Output: false
//	Example 2:
//		Input: arr = [3,5,5]
//		Output: false
//	Example 3:
//		Input: arr = [0,3,2,1]
//		Output: true
//	Constraints:
//		1 <= arr.length <= 104
//		0 <= arr[i] <= 104

func validMountainArray(arr []int) bool {
	l := len(arr)
	if l < 3 {
		return false
	}
	if arr[0] >= arr[1] {
		return false
	}
	i := 1
	// 找到最高点
	for i < l {
		if arr[i-1] == arr[i] {
			return false
		}
		if arr[i-1] > arr[i] {
			break
		}
		i++
	}
	if i >= l {
		return false
	}
	// 最高点为i-1，继续判断i之后的
	i++
	for i < l {
		if arr[i-1] <= arr[i] {
			return false
		}
		i++
	}
	return true
}

func testProblem941() {
	fmt.Println(validMountainArray([]int{2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
	fmt.Println(validMountainArray([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(validMountainArray([]int{1, 2, 3, 2, 1}))
	fmt.Println(validMountainArray([]int{1, 2, 3, 2, 2}))
	fmt.Println(validMountainArray([]int{1, 2, 3, 2, 2, 1}))
	fmt.Println(validMountainArray([]int{3, 2, 1}))
	fmt.Println(validMountainArray([]int{3, 2, 1, 2, 3}))
}
