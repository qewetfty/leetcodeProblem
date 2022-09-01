package bitmap

//Given an integer array nums, return the maximum result of nums[i] XOR nums[j],
// where 0 <= i <= j < n.
//
//
// Example 1:
//
//
//Input: nums = [3,10,5,25,2,8]
//Output: 28
//Explanation: The maximum result is 5 XOR 25 = 28.
//
//
// Example 2:
//
//
//Input: nums = [14,70,53,83,49,91,36,80,92,51,66,70]
//Output: 127
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 2 * 10⁵
// 0 <= nums[i] <= 2³¹ - 1
//
// Related Topics Array Hash Table Bit Manipulation Trie 👍 2760 👎 251

//leetcode submit region begin(Prohibit modification and deletion)
func findMaximumXOR(nums []int) int {
	res := 0
	mask := 0
	for k := 30; k >= 0; k-- {
		// 获得掩码
		mask = mask | (1 << k)
		// 保留前缀
		prefixSet := make(map[int]bool)
		for _, num := range nums {
			prefixSet[num&mask] = true
		}
		// 假设第n位为1
		temp := res | (1 << k)
		for prefix, _ := range prefixSet {
			if _, ok := prefixSet[prefix^temp]; ok {
				res = temp
				break
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
