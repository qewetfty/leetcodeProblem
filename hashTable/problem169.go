package hashTable

import "fmt"

// Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
// You may assume that the array is non-empty and the majority element always exist in the array.
//	Example 1:
//		Input: [3,2,3]
//		Output: 3
//	Example 2:
//		Input: [2,2,1,1,1,2,2]
//		Output: 2

// map method
func majorityElement(nums []int) int {
	numMap := make(map[int]int, len(nums))
	for _, num := range nums {
		if v, ok := numMap[num]; ok {
			v++
			numMap[num] = v
		} else {
			numMap[num] = 1
		}
		if numMap[num] > (len(nums) / 2) {
			return num
		}
	}
	return 0
}

// vote method
func majorityElement2(nums []int) int {
	candidate := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
			if count == 0 {
				candidate = nums[i]
				count = 1
			}
		}
	}
	return candidate
}

func testProblem169() {
	fmt.Println(majorityElement2([]int{8, 8, 7, 7, 7}))
	fmt.Println(majorityElement2([]int{3, 2, 3}))
	fmt.Println(majorityElement2([]int{2, 2, 1, 1, 1, 2, 2}))
}
