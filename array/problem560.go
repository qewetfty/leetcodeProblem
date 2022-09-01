package array

//Given an array of integers nums and an integer k, return the total number of
//continuous subarrays whose sum equals to k.
//
//
// Example 1:
// Input: nums = [1,1,1], k = 2
//Output: 2
// Example 2:
// Input: nums = [1,2,3], k = 3
//Output: 2
//
//
// Constraints:
//
//
// 1 <= nums.length <= 2 * 10⁴
// -1000 <= nums[i] <= 1000
// -10⁷ <= k <= 10⁷
//
// Related Topics Array Hash Table Prefix Sum 👍 10984 👎 353

//leetcode submit region begin(Prohibit modification and deletion)

// 前缀和解法
func subarraySum(nums []int, k int) int {
	l := len(nums)
	// 构造前缀和
	prefixSum := make([]int, l+1)
	for i, num := range nums {
		prefixSum[i+1] = prefixSum[i] + num
	}
	result := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l+1; j++ {
			curSum := prefixSum[j] - prefixSum[i]
			if curSum == k {
				result++
			}
		}
	}
	return result
}

// hashTable+前缀和解法
func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	preSumMap := make(map[int]int)
	preSumMap[0] = 1
	for _, num := range nums {
		pre = pre + num
		if _, ok := preSumMap[pre-k]; ok {
			count = count + preSumMap[pre-k]
		}
		preSumMap[pre]++
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
