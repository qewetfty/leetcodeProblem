package binarySearch

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
//	Example 1:
//		Input: nums1 = [1,3], nums2 = [2]
//		Output: 2.00000
//		Explanation: merged array = [1,2,3] and median is 2.
//	Example 2:
//		Input: nums1 = [1,2], nums2 = [3,4]
//		Output: 2.50000
//		Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
//	Example 3:
//		Input: nums1 = [0,0], nums2 = [0,0]
//		Output: 0.00000
//	Example 4:
//		Input: nums1 = [], nums2 = [1]
//		Output: 1.00000
//	Example 5:
//		Input: nums1 = [2], nums2 = []
//		Output: 2.00000
//	Constraints:
//		nums1.length == m
//		nums2.length == n
//		0 <= m <= 1000
//		0 <= n <= 1000
//		1 <= m + n <= 2000
//		-106 <= nums1[i], nums2[i] <= 106

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(findKthNumber(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return (float64(findKthNumber(nums1, nums2, midIndex1+1)) + float64(findKthNumber(nums1, nums2, midIndex2+1))) / 2
	}
}

func findKthNumber(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return utils.Min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1, newIndex2 := utils.Min(index1+half, len(nums1))-1, utils.Min(index2+half, len(nums2))-1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k = k - (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k = k - (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func testProblem4() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))
}
