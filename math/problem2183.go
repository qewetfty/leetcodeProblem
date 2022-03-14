package math

import "leetcodeProblem/utils"

//Given a 0-indexed integer array nums of length n and an integer k, return the
//number of pairs (i, j) such that:
//
//
// 0 <= i < j <= n - 1 and
// nums[i] * nums[j] is divisible by k.
//
//
//
// Example 1:
//
//
//Input: nums = [1,2,3,4,5], k = 2
//Output: 7
//Explanation:
//The 7 pairs of indices whose corresponding products are divisible by 2 are
//(0, 1), (0, 3), (1, 2), (1, 3), (1, 4), (2, 3), and (3, 4).
//Their products are 2, 4, 6, 8, 10, 12, and 20 respectively.
//Other pairs such as (0, 2) and (2, 4) have products 3 and 15 respectively,
//which are not divisible by 2.
//
//
// Example 2:
//
//
//Input: nums = [1,2,3,4], k = 5
//Output: 0
//Explanation: There does not exist any pair of indices whose corresponding
//product is divisible by 5.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i], k <= 10⁵
//
// Related Topics Array Math Number Theory 👍 349 👎 21

//leetcode submit region begin(Prohibit modification and deletion)
func countPairs(nums []int, k int) int64 {
	gcdFreqMap := make(map[int]int)
	for _, num := range nums {
		gcdFreqMap[utils.Gcd(num, k)]++
	}
	var result int64
	for gcd1, num1 := range gcdFreqMap {
		for gcd2, num2 := range gcdFreqMap {
			if gcd1*gcd2%k == 0 {
				result = result + int64(num1*num2)
			}
		}
	}
	for _, num := range nums {
		if num*num%k == 0 {
			result--
		}
	}
	return result / 2
}
