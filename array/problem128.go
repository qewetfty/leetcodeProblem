package array

import "fmt"

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.
//	Example 1:
//		Input: nums = [100,4,200,1,3,2]
//		Output: 4
//		Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
//	Example 2:
//		Input: nums = [0,3,7,2,5,8,4,6,0,1]
//		Output: 9
//	Constraints:
//		0 <= nums.length <= 105
//		-109 <= nums[i] <= 109

func longestConsecutive(nums []int) int {
	numberMap := make(map[int]bool)
	for _, num := range nums {
		numberMap[num] = true
	}
	result := 0
	for num := range numberMap {
		if numberMap[num-1] {
			continue
		}
		currentLength, currentNum := 1, num
		for numberMap[currentNum+1] {
			currentLength++
			currentNum++
		}
		if currentLength > result {
			result = currentLength
		}
	}
	return result
}

func testProblem128() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2, 4, 4}))
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2, 1, 2, 3}))
}
