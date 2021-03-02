package hashTable

import "fmt"

// You have a set of integers s, which originally contains all the numbers from 1
// to n. Unfortunately, due to some error, one of the numbers in s got duplicated
// to another number in the set, which results in repetition of one number and
// loss of another number.
// You are given an integer array nums representing the data status of this set after the error.
// Find the number that occurs twice and the number that is missing and return them in the form of an array.
//	Example 1:
//		Input: nums = [1,2,2,4]
//		Output: [2,3]
//	Example 2:
//		Input: nums = [1,1]
//		Output: [1,2]
//	Constraints:
//		2 <= nums.length <= 104
//		1 <= nums[i] <= 104

func findErrorNums(nums []int) []int {
	l, numMap := len(nums), make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	result := make([]int, 2)
	for i := 1; i <= l; i++ {
		if numMap[i] == 2 {
			result[0] = i
		}
		if numMap[i] == 0 {
			result[1] = i
		}
	}
	return result
}

func testProblem645() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
}
