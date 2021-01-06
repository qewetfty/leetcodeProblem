package array

import "fmt"

// Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.
// Find the kth positive integer that is missing from this array.
//	Example 1:
//		Input: arr = [2,3,4,7,11], k = 5
//		Output: 9
//		Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.
//	Example 2:
//		Input: arr = [1,2,3,4], k = 2
//		Output: 6
//		Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.
//	Constraints:
//		1 <= arr.length <= 1000
//		1 <= arr[i] <= 1000
//		1 <= k <= 1000
//		arr[i] < arr[j] for 1 <= i < j <= arr.length

// 二分查找算法
func findKthPositive(arr []int, k int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) >> 1
		if arr[mid]-mid-1 >= k {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return k + lo
}

// 线性遍历算法
func findKthPositive2(arr []int, k int) int {
	for i, num := range arr {
		if num-i-1 >= k {
			return k + i
		}
	}
	return k + len(arr)
}

func findKthPositive3(arr []int, k int) int {
	missingNumber := 0
	for i, num := range arr {
		missingNumber = num - i - 1
		if missingNumber >= k {
			return num - (missingNumber - k) - 1
		}
	}
	return arr[len(arr)-1] + k - missingNumber
}

func testProblem1539() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
}
