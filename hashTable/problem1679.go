package hashTable

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// You are given an integer array nums and an integer k.
// In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
// Return the maximum number of operations you can perform on the array.
//	Example 1:
//		Input: nums = [1,2,3,4], k = 5
//		Output: 2
//		Explanation: Starting with nums = [1,2,3,4]:
//			- Remove numbers 1 and 4, then nums = [2,3]
//			- Remove numbers 2 and 3, then nums = []
//			There are no more pairs that sum up to 5, hence a total of 2 operations.
//	Example 2:
//		Input: nums = [3,1,3,4,3], k = 6
//		Output: 1
//		Explanation: Starting with nums = [3,1,3,4,3]:
//			- Remove the first two 3's, then nums = [1,4,3]
//			There are no more pairs that sum up to 6, hence a total of 1 operation.
//	Constraints:
//		1 <= nums.length <= 105
//		1 <= nums[i] <= 109
//		1 <= k <= 109

// hash表一次遍历解法
func maxOperations(nums []int, k int) int {
	numberMap := make(map[int]int)
	result := 0
	for _, num := range nums {
		x := k - num
		if _, ok := numberMap[x]; ok && numberMap[x] > 0 {
			result++
			numberMap[x]--
			continue
		}
		numberMap[num]++
	}
	return result
}

// hash表两遍遍历解法
func maxOperations2(nums []int, k int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}
	numberCountMap := make(map[int]int)
	for _, num := range nums {
		numberCountMap[num]++
	}
	operation := 0
	for num, count := range numberCountMap {
		if num >= k {
			continue
		}
		anotherNum := k - num
		if anotherNum == num {
			operation = operation + count/2
		} else {
			operation = operation + utils.Min(count, numberCountMap[anotherNum])
		}
		delete(numberCountMap, num)
		delete(numberCountMap, anotherNum)
	}
	return operation
}

// 排序双指针解法
func maxOperations3(nums []int, k int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	result := 0
	for i < j {
		sum := nums[i] + nums[j]
		if sum == k {
			result++
			i++
			j--
		} else if sum < k {
			i++
		} else {
			j--
		}
	}
	return result
}

func testProblem1679() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
}
