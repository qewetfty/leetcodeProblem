package array

import "fmt"

//Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
//Note:
//The number of elements initialized in nums1 and nums2 are m and n respectively.
//You may assume that nums1 has enough space (size that is greater or equal to m + n)
// to hold additional elements from nums2.

func testProblem88() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	position := m + n - 1
	pos1 := m - 1
	pos2 := n - 1
	for pos1 >= 0 && pos2 >= 0 {
		number1 := nums1[pos1]
		number2 := nums2[pos2]
		if number1 < number2 {
			nums1[position] = number2
			pos2--
		} else {
			nums1[position] = number1
			pos1--
		}
		position--
	}
	if pos2 >= 0 {
		for i := 0; i <= pos2; i++ {
			nums1[i] = nums2[i]
		}
	}
}
