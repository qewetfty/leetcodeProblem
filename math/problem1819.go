package math

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// You are given an array nums that consists of positive integers.
// The GCD of a sequence of numbers is defined as the greatest integer that divides all the numbers in the sequence evenly.
// For example, the GCD of the sequence [4,6,16] is 2.
// A subsequence of an array is a sequence that can be formed by removing some elements (possibly none) of the array.
// For example, [2,5,10] is a subsequence of [1,2,1,2,4,1,5,10].
// Return the number of different GCDs among all non-empty subsequences of nums.
//	Example 1:
//		Input: nums = [6,10,3]
//		Output: 5
//		Explanation: The figure shows all the non-empty subsequences and their GCDs.
//		The different GCDs are 6, 10, 3, 2, and 1.
//	Example 2:
//		Input: nums = [5,15,40,5,6]
//		Output: 7
//	Constraints:
//		1 <= nums.length <= 105
//		1 <= nums[i] <= 2 * 105

func countDifferentSubsequenceGCDs(nums []int) int {
	visited := make([]bool, 200001)
	maxNum := -1
	for _, num := range nums {
		visited[num] = true
		maxNum = utils.Max(maxNum, num)
	}
	result := 0
	for i := 1; i <= maxNum; i++ {
		commonGCD := -1
		for j := i; j <= maxNum; j += i {
			if visited[j] {
				if commonGCD == -1 {
					commonGCD = j
				} else {
					commonGCD = gcd(commonGCD, j)
				}
			}
		}
		if commonGCD == i {
			result++
		}
	}
	return result
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}

func testProblem1819() {
	fmt.Println(countDifferentSubsequenceGCDs([]int{6, 10, 3}))
	fmt.Println(countDifferentSubsequenceGCDs([]int{5, 15, 40, 5, 6}))
}
