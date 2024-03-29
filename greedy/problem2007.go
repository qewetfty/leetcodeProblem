package greedy

//An integer array original is transformed into a doubled array changed by
//appending twice the value of every element in original, and then randomly shuffling
//the resulting array.
//
// Given an array changed, return original if changed is a doubled array. If
//changed is not a doubled array, return an empty array. The elements in original
//may be returned in any order.
//
//
// Example 1:
//
//
//Input: changed = [1,3,4,2,6,8]
//Output: [1,3,4]
//Explanation: One possible original array could be [1,3,4]:
//- Twice the value of 1 is 1 * 2 = 2.
//- Twice the value of 3 is 3 * 2 = 6.
//- Twice the value of 4 is 4 * 2 = 8.
//Other original arrays could be [4,3,1] or [3,1,4].
//
//
// Example 2:
//
//
//Input: changed = [6,3,0,1]
//Output: []
//Explanation: changed is not a doubled array.
//
//
// Example 3:
//
//
//Input: changed = [1]
//Output: []
//Explanation: changed is not a doubled array.
//
//
//
// Constraints:
//
//
// 1 <= changed.length <= 10⁵
// 0 <= changed[i] <= 10⁵
//
//
// Related Topics Array Hash Table Greedy Sorting 👍 1191 👎 76

//leetcode submit region begin(Prohibit modification and deletion)
func findOriginalArray(changed []int) []int {
	l := len(changed)
	if l%2 == 1 {
		return []int{}
	}
	mapArr := make([]int, 100001)
	for _, num := range changed {
		mapArr[num]++
	}
	expectedLen := l / 2
	result := make([]int, expectedLen)
	curIndex := 0
	for i := 0; i < 100001; i++ {
		for mapArr[i] > 0 {
			mapArr[i]--
			if i*2 > 100000 || mapArr[i*2] == 0 {
				return []int{}
			}
			mapArr[i*2]--
			result[curIndex] = i
			if curIndex++; curIndex == expectedLen {
				break
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
